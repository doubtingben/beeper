package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/doubtingben/beeper/result"
)

// Check that will be passed to beeping
type Check struct {
	URL      string        `json:"url" binding:"required"`
	Pattern  string        `json:"pattern"`
	Header   string        `json:"header"`
	Insecure bool          `json:"insecure"`
	Timeout  time.Duration `json:"timeout"`
}

func main() {

	check := Check{
		URL:      "https://google.fr",
		Pattern:  "find me",
		Header:   "Server: github.com",
		Insecure: false,
		Timeout:  20,
	}

	client := &http.Client{}
	j, err := json.Marshal(check)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPost,
		"http://localhost:8080/check",
		bytes.NewBuffer(j))

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println("status: ", string(resp.Status))
	fmt.Println("headers:")
	for k, v := range resp.Header {
		fmt.Println(" ", k, ":", v)
	}

	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	var result result.Result
	err = dec.Decode(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

}
