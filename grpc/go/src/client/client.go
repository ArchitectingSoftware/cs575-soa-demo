package main

import (
	"context"
	"fmt"
	"log"
	"pubs-proto"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Publication Client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := pubs.NewPubServiceClient(cc)

	doGetPub(c)
}

func doGetPub(c pubs.PubServiceClient) {
	fmt.Println("Starting to query a publication...")
	req := &pubs.PubRequest{
		PubId: 55,
	}

	res, err := c.GetPub(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Publication RPC: %v", err)
	}
	log.Printf("Response from Pub Server: %v", res.GetPub())
}
