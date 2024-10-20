from flask import Flask, request, jsonify
from selenium import webdriver
from selenium.webdriver.support.ui import WebDriverWait

from transformers import pipeline
from bs4 import BeautifulSoup

app = Flask(__name__)

# Summarization and sentiment pipelines
summarizer = pipeline("summarization", model="facebook/bart-large-cnn")
sentiment_analyzer = pipeline("sentiment-analysis")

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
    extension_url = data['url']

    # Step 1: Crawl the reviews using Selenium
    driver = webdriver.Chrome()  # Setup your Chrome driver
    driver.get(extension_url)
    
    # Extract reviews with BeautifulSoup
    page_source = driver.page_source

    soup = BeautifulSoup(page_source, 'html.parser')
    reviews = soup.find_all('section', class_='T7rvce')
    # Step 2: Analyze reviews
    strengths = []
    weaknesses = []
    
    for review in reviews:
        review_text = review.find('p', class_='fzDEpf').text
        print(review_text)
        sentiment = sentiment_analyzer(review_text)[0]['label']
        
        classification = classify_review(review_text, sentiment)
        if classification == 'Strength':
            strengths.append(review_text)
        elif classification == 'Weakness':
            weaknesses.append(review_text)

    print("strength:", strengths)
    print("weakness:", weaknesses)
    # Summarize strengths and weaknesses
    summarized_strengths = summarizer(" ".join(strengths), max_length=100, min_length=50, do_sample=False)[0]['summary_text']
    summarized_weaknesses = summarizer(" ".join(weaknesses), max_length=100, min_length=50, do_sample=False)[0]['summary_text']

    driver.quit()

    # Step 3: Return the results
    return jsonify({
        'strengths': summarized_strengths,
        'weaknesses': summarized_weaknesses,
        'strengths_arr': strengths, 
        'weaknesses_arr': weaknesses,
    })

if __name__ == '__main__':
    app.run(debug=True)
