package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    pb "github.com/adarsh-jaiss/assingment/gen/go" 
)

type server struct{}

func (s *server) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
    // Dummy implementation
    return &pb.UserResponse{
        Username: "tonystark",
        Bio:      "Iron man",
    }, nil
}

func (s *server) GetTweets(ctx context.Context, req *pb.TweetsRequest) (*pb.TweetsResponse, error) {
    // Dummy implementation
    return &pb.TweetsResponse{
        Tweets: []string{"tweet1", "tweet2", "tweet3"},
    }, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    srv := grpc.NewServer()
    pb.RegisterTwitterServiceServer(srv, &server{})
    log.Println("Server listening on port 50051")
    if err := srv.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
