package main

import (
	"context"
	pb "github.com/allexysleeps/dns/dns-proto"
	database "github.com/allexysleeps/dns/domain-address/db"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = ":50051"

var db = database.Create()

type DomainAddressServer struct {
	pb.UnimplementedDomainAddressServer
}

func (s *DomainAddressServer) Save(ctx context.Context, in *pb.NewDomainAddress)(*empty.Empty, error) {
	err := db.Set(in.Domain, in.Address)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *DomainAddressServer) Get(ctx context.Context, in *pb.Domain) (*pb.Address, error) {
	addr, err := db.Get(in.Domain)
	if err != nil {
		return nil, nil
	}
	return &pb.Address{ Address: addr }, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
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