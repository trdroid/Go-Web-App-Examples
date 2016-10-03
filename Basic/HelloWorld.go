package main

import (
  "fmt"
  "net/http"
  "time"
)

const (
  Port = ":8999"
)

func staticFileHandler(writer http.ResponseWriter, request *http.Request) {
  http.ServeFile(writer, request, "static/sample.html")
}

func dynamicRequestHandler(writer http.ResponseWriter, request *http.Request) {
  response := "Hello World! Current Time is " + time.Now().String()
  fmt.Fprintln(writer, response)
}

func main() {
  http.HandleFunc("/static", staticFileHandler)
  http.HandleFunc("/", dynamicRequestHandler)
  http.ListenAndServe(Port, nil)
}
