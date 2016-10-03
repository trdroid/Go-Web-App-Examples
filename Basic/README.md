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

### Browser snapshots

![](_misc/serving%20static%20file.PNG)

![](_misc/serving%20dynamic%20content.PNG)

![](_misc/serving%20unsupported%20route.PNG)

### Request Handlers

The request handler passed as a second argument to *http.HandleFunc()* should have the following prototype

```
func(http.ResponseWriter, *http.Request)
```

Using a handler that does not have the prototype mentioned above results in an error, as shown in the example below:

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

func incorrectHandler() {                           <----------------
}

func main() {
  http.HandleFunc("/static", staticFileHandler)
  http.HandleFunc("/", dynamicRequestHandler)
  http.HandleFunc("/incorrect", incorrectHandler)   <----------------
  http.ListenAndServe(Port, nil)
}
```

```sh
$ go run HelloWorld.go
# command-line-arguments
.\HelloWorld.go:28: cannot use incorrectHandler (type func()) as type func(http.ResponseWriter, *http.Request) in argument to http.HandleFunc
```
