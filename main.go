package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
)

func main() {
  fmt.Println("garden")

  response, err := http.Get("https://trefle.io/api/v1/plants?token=O65c8cMk7fhAmpwC4pS-PX8YjSZTsDapKsE3tG6MKO4")

  if err != nil {
    fmt.Println("error requesting. %s\n", err)
  } else {
    data, _ := ioutil.ReadAll(response.Body)
    fmt.Println(string(data))
  }
}
