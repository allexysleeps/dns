package main

import (
	"context"
	pb "github.com/allexysleeps/dns/dns-proto"
	database "github.com/allexysleeps/dns/domain-address/db"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = ":50051"

var db = database.Create()

type DomainAddressServer struct {
	pb.UnimplementedDomainAddressServer
}

func (s *DomainAddressServer) Save(ctx context.Context, in *pb.NewDomainAddress)(*pb.Address, error) {
	err := db.Set(in.Domain, in.Address)
	if err != nil {
		return nil, err
	}
	log.Printf("%s address saved", in.Domain)
	return &pb.Address{Address: in.Address}, nil
}

func (s *DomainAddressServer) Get(ctx context.Context, in *pb.Domain) (*pb.Address, error) {
	addr, err := db.Get(in.Domain)
	if err != nil {
		return nil, nil
	}
	return &pb.Address{ Address: addr }, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDomainAddressServer(s, &DomainAddressServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}