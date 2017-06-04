package main

import (
  "net/http"
  "net/url"
)

import "io/ioutil"
import "fmt"
import "strings"

func keepLines(s string, n int) string {
  result := strings.Join(strings.Split(s, "\n")[:n], "\n")
  return strings.Replace(result, "\r", "", -1)
}

func main() {

  resp, err := http.PostForm("http://localhost:8080/check", url.Values{"url": {"https://google.fr"}})
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)

  fmt.Println("status: ", string(resp.Status))
  fmt.Println("headers:")
  for k, v := range resp.Header {
    fmt.Println(" ", k, ":", v)
  }
  fmt.Println("post:\n", keepLines(string(body), 3))

}
