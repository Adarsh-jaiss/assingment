package main

import (
    "context"
    "log"

    "google.golang.org/grpc"
    pb "github.com/adarsh-jaiss/assingment/gen/go"  
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    client := pb.NewTwitterServiceClient(conn)

    // Calling GetUser
    userReq := &pb.UserRequest{UserId: "123"}
    userResp, err := client.GetUser(context.Background(), userReq)
    if err != nil {
        log.Fatalf("Error calling GetUser: %v", err)
    }
    log.Printf("User Response: %v", userResp)

    // Call GetTweets
    tweetsReq := &pb.TweetsRequest{Hashtag: "#intern"}
    tweetsResp, err := client.GetTweets(context.Background(), tweetsReq)
    if err != nil {
        log.Fatalf("Error calling GetTweets: %v", err)
    }
    log.Printf("Tweets Response: %v", tweetsResp)
}
