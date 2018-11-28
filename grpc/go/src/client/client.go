package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"pubs-proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	tls = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
)

func main() {
	fmt.Println("Publication Client")
	flag.Parse()

	var opts []grpc.DialOption
	//now the server options
	if *tls {
		creds, err := credentials.NewClientTLSFromFile("../ssl/server.crt", "")

		if err != nil {
			log.Fatalf("Failed to TLS files: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	cc, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := pubs.NewPubServiceClient(cc)

	doStreamAllPubs(c)
}

func doGetPub(c pubs.PubServiceClient) {
	fmt.Println("Starting to query a publication...")
	req := &pubs.PubRequest{
		PubId: 5,
	}

	res, err := c.GetPub(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Publication RPC: %v", err)
	}
	log.Printf("Response from Pub Server: %v", res.GetPub())
}

func doGetAllPubs(c pubs.PubServiceClient) {
	fmt.Println("Starting to query for ALL publications...")
	req := &pubs.EmptyRequest{}

	res, err := c.GetAllPubs(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Publication RPC: %v", err)
	}

	for i, p := range res.GetPubs() {
		log.Printf("%d: Pub Title: %s \n", i, p.GetTitle())
	}
}

func doStreamAllPubs(c pubs.PubServiceClient) {
	fmt.Println("Starting to query for ALL publications...")
	req := &pubs.EmptyRequest{}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := c.StreamPubs(ctx, req)
	if err != nil {
		log.Fatalf("Fatal error while calling for Publication RPC Stream: %v", err)
	}

	for {
		p, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Fatal error while receiving from the stream: %v", err)
		}
		log.Printf("|-- Pub Title --> %s \n", p.GetTitle())
	}
}
