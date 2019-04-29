package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "./cats"
)

const (
	port = ":50051"
)

// server is used to implement cats.CatServiceServer.
type server struct{
	catsList pb.CatList
}

// SayHello implements helloworld.GreeterServer
// func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
// 	log.Printf("Received: %v", in.Name)
// 	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
// }

func (s *server) List(ctx context.Context, req *pb.Empty) (*pb.CatList, error) {
	return &s.catsList, nil
}

func (s *server) Delete(ctx context.Context, req *pb.CatIdRequest) (*pb.Empty, error) {
	return nil, nil
}

func (s *server) Insert(ctx context.Context, req *pb.Cat) (*pb.Empty, error) {
	return nil, nil
}

func (s *server) Get(ctx context.Context, req *pb.CatIdRequest) (*pb.Cat, error) {
	return nil, nil
}

func (s *server) Watch(req *pb.Empty, srv pb.CatService_WatchServer) error {
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	catServer := server{};

	s := grpc.NewServer()
	pb.RegisterCatServiceServer(s, &catServer)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
