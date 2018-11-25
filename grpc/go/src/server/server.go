package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"pubs-proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

var pubCache []PubDbRecord = nil

func (*server) GetPub(ctx context.Context, req *pubs.PubRequest) (*pubs.PubResponse, error) {
	fmt.Printf("Get function was invoked with %v\n", req)
	pubID := int(req.GetPubId())

	if pubCache == nil {
		log.Fatalf("Publication Cache is Not Loaded")
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: Publication Cache was Not Loaded"),
		)
	}

	if (pubID < 1) || pubID >= len(pubCache) {
		log.Printf("Request is out of range")
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Internal error: Publication ID is not Found"),
		)
	}

	tPub := pubCache[pubID-1]
	pResp := pubs.Pub{
		Id:       int32(tPub.ID),
		Title:    tPub.Title,
		Cite:     tPub.Cite,
		Abstract: tPub.Abstract,
		Link:     tPub.Link,
		Slides:   tPub.Slides,
	}

	res := &pubs.PubResponse{
		Pub: &pResp,
	}
	return res, nil
}

func main() {
	fmt.Println("Publication Server")

	pc, err := LoadPubs()
	if err != nil {
		log.Fatalf("Failed to load publication cache: %v", err)
	}
	//set the cache
	pubCache = pc

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pubs.RegisterPubServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
