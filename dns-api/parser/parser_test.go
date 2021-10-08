package parser

import (
	"strings"
	"testing"
)

func TestParseSuccess(t *testing.T) {
	json := `
	[
		{
			"domain": "tst.dom",
			"port": 9000,
			"ip": "123.100.1.10"
		},
		{
			"domain": "domn.com",
			"port": 10010,
			"ip": "10.120.12.100"
		},
	]
	`
	want := []Record{
		{
			Domain: "tst.dom",
			Ip:     "123.100.1.10",
			Port:   9000,
		},
		{
			Domain: "domn.com",
			Ip:     "10.120.12.100",
			Port:   10010,
		},
	}

	records := make(chan Record)
	go Parse(strings.NewReader(json), records)

	for i := 0; i < len(want); i++ {
		r := <-records
		if r != want[i] {
			t.Errorf("incorrectly parsed chunk want: %v, got: %v", want[i], r)
		}
	}
	if _, ok := <-records; ok {
		t.Errorf("channel has not been closed")
	}
}
