*Basic/HelloWorld.go*

```go
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
```

*Basic/static/sample.html*

```html
<!DOCTYPE html>
<html>
<head>
  <title>Hello World!</title>
</head>
<body>
  <h1>Hello World! I am a static file</h1>
</body>
</html>
```
