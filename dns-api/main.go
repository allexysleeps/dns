package main

import (
	"fmt"
	"github.com/allexysleeps/dns/dns-api/router"
	pb "github.com/allexysleeps/dns/dns-proto"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

const address = "dns-storage:50051"

func main() {
	fmt.Println("Starting the server")
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewDnsStorageClient(conn)

	r := router.CreateRouter(client)
	http.Handle("/", r)

	log.Println("Server is running on: localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
