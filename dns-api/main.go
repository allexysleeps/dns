package main

import (
  "context"
  "log"
  "net/http"
  "time"
)

const address = "localhost:50051"


func main() {
  c := createRpcClient()
  
  ctx, cancel := context.WithTimeout(context.Background(), time.Second)
  defer cancel()
  
  r := createRouter(c, ctx)
  http.Handle("/", r)

  log.Println("Server is running on: localhost:8080")
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}