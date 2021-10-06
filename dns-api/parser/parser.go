package parser

import (
	"encoding/json"
	"errors"
	"io"
)

type Record struct {
	Domain, Ip string
	Port       int
}

var parsingError = errors.New("couldn't parse JSON")

func Parse(input io.Reader, records chan<- Record) error {
	defer close(records)
	dec := json.NewDecoder(input)
	_, err := dec.Token()
	if err != nil {
		return parsingError
	}
	for dec.More() {
		var r Record
		err := dec.Decode(&r)
		if err != nil {
			return parsingError
		}
		records <- r
	}

	_, err = dec.Token()
	if err != nil {
		return parsingError
	}
	return nil
}
