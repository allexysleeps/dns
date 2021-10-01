package jsonStreamParse

import (
  "dns-api/shared"
  "encoding/json"
  "io"
  "log"
)

func Parse(input *io.Reader, records chan<- shared.Record) {
  dec := json.NewDecoder(*input)
  _, err := dec.Token()
  if err != nil {
    log.Fatal(err)
  }
  for dec.More() {
    var r shared.Record
    err := dec.Decode(&r)
    if err != nil {
      log.Fatal(err)
    }
    records <- r
  }
  close(records)
  _, err = dec.Token()
  if err != nil {
    log.Fatal(err)
  }
}