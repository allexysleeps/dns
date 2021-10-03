package main

import (
  "context"
  "github.com/allexysleeps/dns/dns-api/parser"
  "github.com/allexysleeps/dns/dns-api/shared"
  pb "github.com/allexysleeps/dns/dns-proto"
  "github.com/gorilla/mux"
  "google.golang.org/grpc"
  "io"
  "log"
  "net/http"
  "net/url"
  "strconv"
)

func submitRecords(client pb.DomainAddressClient, ctx context.Context) http.HandlerFunc {
  return func (w http.ResponseWriter, r *http.Request) {
    file, _, err := r.FormFile("file")
    if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      return
    }
    
    reader := io.Reader(file)
    
    records := make(chan shared.Record)
    
    go parser.Parse(&reader, records)
    
    for r := range records {
      addr := r.Ip + ":" + strconv.Itoa(r.Port)
      _, err := client.Save(ctx, &pb.NewDomainAddress{Domain: r.Domain, Address: addr})
      if err != nil {
        log.Printf("couldnt save: %s:%s", r.Domain, addr)
      }
    }
  }
}

func getRecord(client pb.DomainAddressClient, ctx context.Context) http.HandlerFunc {
  return func (w http.ResponseWriter, r *http.Request) {
    q, err := url.ParseQuery(r.URL.RawQuery)
    if err != nil || !q.Has("domain") {
      w.WriteHeader(http.StatusBadRequest)
      return
    }
    addr, err := client.Get(ctx, &pb.Domain{Domain: q.Get("domain")})
    if err != nil {
      w.WriteHeader(http.StatusNotFound)
      return
    }
    w.Write([]byte(addr.Address))
  }
}

func createRouter(c pb.DomainAddressClient, ctx context.Context) *mux.Router {
  r := mux.NewRouter()
  
  r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
  }).Methods(http.MethodGet)
  r.HandleFunc("/table", submitRecords(c, ctx)).Methods(http.MethodPost)
  r.HandleFunc("/address", getRecord(c, ctx)).Methods(http.MethodGet)
  
  return r
}

func createRpcClient() pb.DomainAddressClient {
  conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
  
  if err != nil {
    log.Fatalf("did not connect: %v", err)
  }
  defer conn.Close()
  c := pb.NewDomainAddressClient(conn)
  
  return c
}