package main

import (
    "net/http"
    "fmt"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello %s", r.URL.Path[1:])
}

func main() {
    fmt.Println("listen localhost:3333")
    http.HandleFunc("/", handler)
    http.ListenAndServe(":3333", nil)

}
