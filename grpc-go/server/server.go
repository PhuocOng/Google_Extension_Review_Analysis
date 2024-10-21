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
	return &pb.RecommendationResponse{Recommendations: "Example Testing Text for postman"}, nil
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
