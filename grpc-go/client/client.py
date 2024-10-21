from flask import Flask, request, jsonify
import grpc
import recommendations_pb2
import recommendations_pb2_grpc
from flask import Flask, request, jsonify
from selenium import webdriver
import time
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from bs4 import BeautifulSoup
from transformers import pipeline
from flask_cors import CORS
import redis
from datetime import datetime, timedelta

app = Flask(__name__)
CORS(app)

# Summarization and sentiment pipelines
summarizer = pipeline("summarization", model="facebook/bart-large-cnn")
sentiment_analyzer = pipeline("sentiment-analysis")

redis_client = redis.Redis(
  host='redis-10992.c9.us-east-1-2.ec2.redns.redis-cloud.com',
  port=10992,
  password='AQhSUkd75maPURHJl7YikDalJuFTH2bk')

def get_ttl_for_today():
    now = datetime.now()
    # Calculate the time left until midnight
    midnight = datetime.combine(now + timedelta(days=1), datetime.min.time())
    ttl = (midnight - now).seconds  # Time in seconds until midnight
    return ttl

def get_recommendations(strengths, weaknesses):
    # Connect to the gRPC server running on the Golang backend
    channel = grpc.insecure_channel('localhost:50051')
    stub = recommendations_pb2_grpc.RecommendationServiceStub(channel)

    # Send the request with strengths and weaknesses
    request = recommendations_pb2.RecommendationRequest(
        strengths=strengths,
        weaknesses=weaknesses
    )
    response = stub.GetRecommendations(request)

    return response.recommendations

# Keywords for strength/weakness detection
bug_keywords = [
'crash', 'error', 'freeze', 'slow', 'problem', 'issue', 'not working', 'missing', 'fail', 'broken', 
'stuck', 'unresponsive', 'lag', 'bug', 'disconnect', 'failing', 'glitch', 'does not load', 'overheating', 
'screen flickering', 'unable to install', 'unexpected behavior', 'restart', 'does not open', 
'incompatibility', 'corrupt', 'out of memory', 'malfunction', 'hang', 'black screen', 'file corrupted', 
'delayed response', 'UI freezes', 'service outage', 'won’t connect', 'time out', 'conflict', 'data loss', 
'installation failed', 'unstable', 'crashing on start', "can't access", 'unintuitive', 'no support', 
'overload', 'no sound', "files won't upload", "files won't download", 'inconsistent performance', 
'no sound', 'not', 'no', 'limited', "can't", "but"
]

strength_keywords = [
'works well', 'fast', 'easy', 'seamless', 'great', 'useful', 'smooth', 'helpful', 'efficient', 
'intuitive', 'reliable', 'well-designed', 'user-friendly', 'lightweight', 'stable', 'responsive', 
'consistent', 'secure', 'high quality', 'excellent', 'performance', 'good interface', 'customizable', 
'feature-rich', 'robust', 'multifunctional', 'speedy', 'clear instructions', 'no lag', 'quick setup', 
'time-saving', 'productive', 'well-supported', 'versatile', 'easily accessible', 'clear design', 
'good customer support', 'multi-platform', 'rich features', 'file upload works', 'great for collaboration', 
'reliable connection', 'minimal issues', 'effective', 'strong functionality', 'saves time', 'streamlined'
]

# Function to classify reviews
def classify_review(summary_text, sentiment):
    # Detect if it's a strength or weakness
    bug_detected = any(kw in summary_text for kw in bug_keywords)
    strength_detected = any(kw in summary_text for kw in strength_keywords)

    if sentiment == 'NEGATIVE' and bug_detected:
        return 'Weakness'
    elif sentiment == 'POSITIVE' and strength_detected:
        return 'Strength'
    else:
        return 'Neutral'

@app.route('/api/analyze-extension', methods=['POST'])
def analyze_extension():
    data = request.json
    extension_url = data['url'] + "/reviews"

    cached_result = redis_client.get(extension_url)
    if cached_result:
        return jsonify(eval(cached_result))  # Return cached result

    # Step 1: Crawl the reviews using Selenium
    driver = webdriver.Chrome()  # Setup your Chrome driver
    driver.get(extension_url)

    wait = WebDriverWait(driver, 10)

# Click the "Load more" button to get more reviews
    try:
        for _ in range(5):
            # Wait until the "Load more" button is clickable
            load_more_button = wait.until(EC.element_to_be_clickable((By.XPATH, "//span[contains(text(),'Load more')]")))
            load_more_button.click()
            time.sleep(2)  # Briefly wait for reviews to load after clicking
    except Exception as e:
        print(f"No more 'Load more' button or an error occurred: {e}")
    
    # Extract reviews with BeautifulSoup
    page_source = driver.page_source

    soup = BeautifulSoup(page_source, 'html.parser')
    reviews = soup.find_all('section', class_='T7rvce')
    # Step 2: Analyze reviews
    strengths = []
    weaknesses = []
    
    for review in reviews:
        review_text = review.find('p', class_='fzDEpf').text
        sentiment = sentiment_analyzer(review_text)[0]['label']
        
        classification = classify_review(review_text, sentiment)
        if classification == 'Strength':
            strengths.append(review_text)
        elif classification == 'Weakness':
            weaknesses.append(review_text)

    # Summarize strengths and weaknesses
    summarized_strengths = summarizer(" ".join(strengths), max_length=100, min_length=50, do_sample=False)[0]['summary_text']
    summarized_weaknesses = summarizer(" ".join(weaknesses), max_length=100, min_length=50, do_sample=False)[0]['summary_text']
    recommendations = get_recommendations(" ".join(strengths), " ".join(weaknesses))

    driver.quit()

    # Step 3: Return the results
    result = {
        'strengths': summarized_strengths,
        'weaknesses': summarized_weaknesses,
        'strengths_arr': strengths,
        'weaknesses_arr': weaknesses,
        'recommendations': recommendations,
    }
    
    # Cache the result in Redis with TTL until midnight
    ttl = get_ttl_for_today()  # Get TTL in seconds until midnight
    redis_client.setex(extension_url, ttl, str(result))  # Cache the result 
    return jsonify(result)

if __name__ == '__main__':
    app.run(debug=True)

# Initialize the Flask application
app = Flask(__name__)

# Function to call the Golang gRPC service and get recommendations


# # Example usage in a Flask route after crawling and analyzing the reviews
# @app.route('/api/analyze-extension', methods=['POST'])
# def analyze_extension():
#     # Example data; in practice, you would get these from crawling and summarizing reviews
#     data = request.json
#     strengths = data.get("strengths", "fast, easy to use")
#     weaknesses = data.get("weaknesses", "crashes often, hard to install")

#     # Call gRPC service for recommendations

#     # Return the analysis results along with recommendations
#     return jsonify({
#         "strengths": strengths,
#         "weaknesses": weaknesses,
#         "recommendations": recommendations
#     })

# Run the Flask app
if __name__ == '__main__':
    app.run(debug=True)
