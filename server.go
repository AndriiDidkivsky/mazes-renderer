package main

import (
	"./router"
	"fmt"
	"net/http"
)

// func indexHandler(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "Hello %s", r.URL.Path[1:])
// }

func main() {
	router := router.New()
	fmt.Println("listen localhost:3333")
	// http.HandleFunc("/", handler)
	http.ListenAndServe(":3333", router)

}
