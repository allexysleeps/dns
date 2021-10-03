package main

import (
  "bufio"
  "encoding/json"
  "fmt"
  "log"
  "math/rand"
  "os"
  "strconv"
  "strings"
  "time"
)

const targetDir = "../generated"

type record struct {
  Domain string `json:"domain"`
  Ip     string `json:"ip"`
  Port   int    `json:"port"`
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")
var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateString(size int) string {
  var result string
  for i := 0; i < size; i++ {
    r := letterRunes[random.Intn(len(letterRunes) - 1)]
    result += string(r)
  }
  return result
}

func generateDomain() string {
  nameSize := random.Intn(13) + 2
  zoneSize := random.Intn(3) + 2
  return fmt.Sprintf("%s.%s", generateString(nameSize), generateString(zoneSize))
}

func generateIP() string {
  ip := make([]string, 0, 4)
  for i := 0; i < 4; i++ {
    ip = append(ip, strconv.Itoa(random.Intn(255)))
  }
  return strings.Join(ip, ".")
}

func generateRecord(n int) record {
  return record {
    Domain: generateDomain(),
    Ip: generateIP(),
    Port: random.Intn(65535),
  }
}

func check(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func main() {
    nRecords, err := strconv.Atoi(os.Args[1])
    if err != nil {
        log.Print("could get the file size")
        log.Fatal(err)
    }
    if _, err := os.Stat(targetDir); os.IsNotExist(err) {
        if err := os.Mkdir(targetDir, 0755); err != nil {
            log.Fatal(err)
        }
    }

    fName := fmt.Sprintf("%s/%d-lines.json", targetDir, nRecords)

    file, err := os.Create(fName)
    defer file.Close()
    check(err)
    w := bufio.NewWriter(file)
    _, err = w.WriteRune('[')
    check(err)
    
    for i := 0; i < nRecords; i++ {
      jBytes, err := json.Marshal(generateRecord(i))
      fmt.Printf("\r%d/%d", i + 1, nRecords)
      check(err)
      _, err = w.Write(jBytes)
      check(err)
      if i != nRecords - 1 {
        _, err = w.WriteRune(',')
        check(err)
      }
    }
    
    _, err = w.WriteRune(']')
    check(err)
    w.Flush()

    f, err := os.Stat(fName)
    if err != nil {
        fmt.Printf("\nDONE.")
    } else {
        fmt.Printf("\nDONE. File size: %dMB", f.Size() / 1e6)
    }
}
