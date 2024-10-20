from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
import time
from bs4 import BeautifulSoup
from transformers import pipeline
import re

#Step 1: Crawling Comments
# Path to your ChromeDriver
chrome_driver_path = "C:/Users/ASUS/OneDrive/Desktop_old/Tech/Random/chromedriver.exe"

# Setup the Chrome WebDriver
service = Service(chrome_driver_path)
driver = webdriver.Chrome(service=service)

# Open the Chrome Web Store extension reviews page
url = "https://chromewebstore.google.com/detail/chrome-remote-desktop/inomeogfingihgjfjlpeplalcfajhgai/reviews"
driver.get(url)

# Wait for the page to load
wait = WebDriverWait(driver, 10)

# Click the "Load more" button to get more reviews
try:
    for _ in range(1):
        # Wait until the "Load more" button is clickable
        load_more_button = wait.until(EC.element_to_be_clickable((By.XPATH, "//span[contains(text(),'Load more')]")))
        load_more_button.click()
        time.sleep(2)  # Briefly wait for reviews to load after clicking
except Exception as e:
    print(f"No more 'Load more' button or an error occurred: {e}")

# Extract the page source with all reviews loaded
page_source = driver.page_source

# Close the browser
driver.quit()

# You can now use BeautifulSoup to parse the `page_source`


soup = BeautifulSoup(page_source, 'html.parser')

# Example: Extract all review texts
reviews = soup.find_all('section', class_='T7rvce')  # You may need to update the class based on the HTML structure
reviews_text = []
print(reviews)
for i in range(len(reviews)):
    review = reviews[i]
    review_text = review.find('p', class_='fzDEpf').text if review.find('p', class_='fzDEpf') else ""
    print(f"{i}. {review_text}", end = "\n ___________________ \n")
    reviews_text.append(review_text)


print(reviews_text)


def detect_bug_or_strength(text, sentiment):
    # Define bug and strength keywords
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

    # Check for bug-related keywords in the review text
    bug_detected = any(re.search(rf"\b{kw}\b", text, re.IGNORECASE) for kw in bug_keywords)
    strength_detected = any(re.search(rf"\b{kw}\b", text, re.IGNORECASE) for kw in strength_keywords)

    # Classify the review based on sentiment and keywords
    if sentiment[0]['label'] == 'NEGATIVE' and bug_detected:
        return 'Bug Detected', text
    elif sentiment[0]['label'] == 'POSITIVE' and strength_detected:
        return 'Strength Detected', text
    else:
        return 'Neutral or No Significant Feedback', text
    
sentiment_analyzer = pipeline("sentiment-analysis")
for r in reviews_text:
    sentiment = sentiment_analyzer(r)
    print(f"{r} =======> {sentiment[0]['label']}")
    print(detect_bug_or_strength(r, sentiment))
reviews_text = " ".join(reviews_text)

summarizer = pipeline("summarization", model="facebook/bart-large-cnn")

summary = summarizer(reviews_text, max_length=150, min_length=25, do_sample=False)
print("Summary_text is: ")
print(summary[0]['summary_text'])


