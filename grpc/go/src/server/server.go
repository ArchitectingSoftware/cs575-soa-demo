package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"pubs-proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

var (
	tls = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
)

type server struct {
	cache []PubDbRecord
}

//var pubCache []PubDbRecord = nil

func (s *server) GetPub(ctx context.Context, req *pubs.PubRequest) (*pubs.PubResponse, error) {
	fmt.Printf("Get function was invoked with pubID = %v\n", req)
	pubID := int(req.GetPubId())

	if s.cache == nil {
		log.Fatalf("Publication Cache is Not Loaded")
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: Publication Cache was Not Loaded"),
		)
	}

	if (pubID < 1) || pubID >= len(s.cache) {
		log.Printf("Request is out of range")
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Internal error: Publication ID is not Found"),
		)
	}

	tPub := s.cache[pubID-1]
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

func (s *server) GetAllPubs(ctx context.Context, req *pubs.EmptyRequest) (*pubs.PubsResponse, error) {
	fmt.Printf("GetAllPubs function was invoked\n")
	if s.cache == nil {
		log.Fatalf("Publication Cache is Not Loaded")
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: Publication Cache was Not Loaded"),
		)
	}

	rspData := make([]*pubs.Pub, len(s.cache))
	for i, p := range s.cache {
		rspData[i] = &pubs.Pub{
			Id:       int32(p.ID),
			Title:    p.Title,
			Cite:     p.Cite,
			Abstract: p.Abstract,
			Link:     p.Link,
			Slides:   p.Slides,
		}

	}

	res := &pubs.PubsResponse{
		Pubs: rspData,
	}
	return res, nil
}

func (s *server) StreamPubs(req *pubs.EmptyRequest, stream pubs.PubService_StreamPubsServer) error {
	fmt.Printf("StreamPubs function was invoked\n")
	if s.cache == nil {
		log.Fatalf("Publication Cache is Not Loaded")
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: Publication Cache was Not Loaded"),
		)
	}

	for _, p := range s.cache {
		rspData := &pubs.Pub{
			Id:       int32(p.ID),
			Title:    p.Title,
			Cite:     p.Cite,
			Abstract: p.Abstract,
			Link:     p.Link,
			Slides:   p.Slides,
		}

		if err := stream.Send(rspData); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	log.Printf("Starting Publication Server\n")
	flag.Parse()

	pc, err := LoadPubs()
	if err != nil {
		log.Fatalf("Failed to load publication cache: %v", err)
	}

	//set the cache
	serverConfig := &server{
		cache: pc,
	}

	var opts []grpc.ServerOption
	//now the server options
	if *tls {
		creds, err := credentials.NewServerTLSFromFile("../ssl/server.crt", "../ssl/server.key")

		if err != nil {
			log.Fatalf("Failed to TLS files: %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer(opts...)
	pubs.RegisterPubServiceServer(s, serverConfig)

	log.Printf("Ready for Requests...\n")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
