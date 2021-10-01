package main

import (
  "dns-api/db"
  "dns-api/jsonStreamParse"
  "dns-api/shared"
  "io"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "net/url"
)



func submitTable(w http.ResponseWriter, r *http.Request) {
  file, _, err := r.FormFile("file")
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  reader := io.Reader(file)

  records := make(chan shared.Record)

  go jsonStreamParse.Parse(&reader, records)
  go db.Save(records)
}

func getRecord(w http.ResponseWriter, r *http.Request) {
  q, err := url.ParseQuery(r.URL.RawQuery)
  if err != nil || !q.Has("domain") {
    w.WriteHeader(http.StatusBadRequest)
    return
  }
  addr, ok := db.Get(q.Get("domain"))
  if !ok {
    w.WriteHeader(http.StatusNotFound)
    return
  }
  w.Write([]byte(addr))
}

func main() {
  r := mux.NewRouter()

  r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
  }).Methods(http.MethodGet)
  r.HandleFunc("/table", submitTable).Methods(http.MethodPost)
  r.HandleFunc("/address", getRecord).Methods(http.MethodGet)

  http.Handle("/", r)

  log.Println("Server is running on: localhost:8080")
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}