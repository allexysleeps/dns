package db

import (
  "dns-api/shared"
  "fmt"
)

var db = make(map[string]string)

func Save(records <-chan shared.Record) {
  for r := range records {
    db[r.Domain] = fmt.Sprintf("%s:%d", r.Ip, r.Port)
  }
}

func Get(domain string) (string, bool) {
  addr, ok := db[domain]
  if !ok {
    return "", false
  }
  return addr, true
}