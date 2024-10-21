from flask import Flask, request, jsonify
import grpc
import recommendations_pb2
import recommendations_pb2_grpc

# Initialize the Flask application
app = Flask(__name__)

# Function to call the Golang gRPC service and get recommendations
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

# Example usage in a Flask route after crawling and analyzing the reviews
@app.route('/api/analyze-extension', methods=['POST'])
def analyze_extension():
    # Example data; in practice, you would get these from crawling and summarizing reviews
    data = request.json
    strengths = data.get("strengths", "fast, easy to use")
    weaknesses = data.get("weaknesses", "crashes often, hard to install")

    # Call gRPC service for recommendations
    recommendations = get_recommendations(strengths, weaknesses)

    # Return the analysis results along with recommendations
    return jsonify({
        "strengths": strengths,
        "weaknesses": weaknesses,
        "recommendations": recommendations
    })

# Run the Flask app
if __name__ == '__main__':
    app.run(debug=True)
