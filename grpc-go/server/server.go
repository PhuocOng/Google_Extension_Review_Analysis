package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	pb "grpc-go/protos" // Import the generated protobuf package
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	"google.golang.org/grpc"
)

// OpenAI API response struct
type OpenAIResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

// GetRecommendations implements the gRPC server method.
func (s *server) GetRecommendations(ctx context.Context, req *pb.RecommendationRequest) (*pb.RecommendationResponse, error) {
	example_text := `Enhance Sync Features: Several users mentioned issues with syncing settings across devices, particularly with the whitelist/blacklist. Implementing a more reliable sync system that remembers user preferences (like the whitelist/blacklist) across different devices would significantly improve the user experience.

	Improve Customization Control: Some users find the app difficult to control. Adding more intuitive customization options or simplifying the current controls can help users feel more in charge of their settings, especially for managing site-specific behaviors.

	Bug Fixes and Stability: Users reported issues with the dark mode not working consistently (e.g., staying in light mode sometimes). Addressing these bugs will ensure a more stable experience for everyone.

	Add More Integrations: There are requests for integration with popular platforms like Google Drive. Expanding the app’s integration capabilities could enhance its usefulness, particularly for power users who use multiple services.

	Documentation and Tutorials: Some users find it difficult to figure out how to make the app work on all websites. Offering clear instructions or a tutorial for first-time users would help reduce the learning curve and improve user satisfaction.

	Selective Website Activation: Users requested an option to deactivate the theme for all websites by default and activate it only for specific sites. Providing this feature would give users more flexibility and control over where the theme is applied.`
	return &pb.RecommendationResponse{Recommendations: example_text}, nil
	apiKey := os.Getenv("OPENAI_API_KEY")
	url := "https://api.openai.com/v1/completions"
	prompt := fmt.Sprintf("Here are the strengths: %s. Here are the weaknesses: %s. Provide solutions.", req.Strengths, req.Weaknesses)

	// Prepare OpenAI API request
	data := map[string]interface{}{
		"model":      "text-davinci-003",
		"prompt":     prompt,
		"max_tokens": 300,
	}

	jsonData, _ := json.Marshal(data)

	// Set up request
	reqOpenAI, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	reqOpenAI.Header.Set("Authorization", "Bearer "+apiKey)
	reqOpenAI.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(reqOpenAI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var openAIResp OpenAIResponse
	if err := json.Unmarshal(body, &openAIResp); err != nil {
		return nil, err
	}

	// Return the recommendation
	return &pb.RecommendationResponse{Recommendations: openAIResp.Choices[0].Text}, nil
}

// Implement the gRPC server
type server struct {
	pb.UnimplementedRecommendationServiceServer
}

func main() {
	// Set up gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterRecommendationServiceServer(s, &server{})

	fmt.Println("gRPC server listening on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
