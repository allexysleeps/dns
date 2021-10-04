package router

import (
  "context"
  "github.com/allexysleeps/dns/dns-api/parser"
  pb "github.com/allexysleeps/dns/dns-proto"
  "github.com/gorilla/mux"
  "io"
  "log"
  "net/http"
  "net/url"
  "strconv"
)

func submitRecords(client pb.DomainAddressClient) http.HandlerFunc {
  return func (w http.ResponseWriter, r *http.Request) {
    file, _, err := r.FormFile("file")
    if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      return
    }
    
    reader := io.Reader(file)
    
    records := make(chan parser.Record)
    
    go func() {
      err := parser.Parse(reader, records)
      if err != nil {
        w.WriteHeader(http.StatusBadRequest)
      }
    }()
    
    for record := range records {
      addr := record.Ip + ":" + strconv.Itoa(record.Port)
      _, err := client.Save(r.Context(), &pb.NewDomainAddress{Domain: record.Domain, Address: addr})
      if err != nil {
        log.Print(err)
        log.Printf("couldnt save: %s:%s", record.Domain, addr)
      }
    }
  }
}

func getRecord(client pb.DomainAddressClient) http.HandlerFunc {
  return func (w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
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

func CreateRouter(c pb.DomainAddressClient) *mux.Router {
  r := mux.NewRouter()
  
  r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
  }).Methods(http.MethodGet)
  r.HandleFunc("/table", submitRecords(c)).Methods(http.MethodPost)
  r.HandleFunc("/address", getRecord(c)).Methods(http.MethodGet)
  
  return r
}