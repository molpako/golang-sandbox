package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	url = "http://localhost:8000/patterns.json"
)

// JSONData is a patterns file for output
type JSONData struct {
	Patterns []string `json:"patterns"`
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

func main() {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	var jd JSONData
	if err := decodeBody(resp, &jd); err != nil {
		panic(err)
	}

	s := strings.Join(jd.Patterns, "\n")
	r := strings.NewReader(s)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
