package main

import (
	"./router"
	"encoding/json"
	"fmt"
	"github.com/AndriiDidkivsky/go-algorithms/btree_maze"
	"net/http"
)

func btreeMaze(w http.ResponseWriter, r *http.Request) {
	maze := btree_maze.GenerateMaze(5, 5)
	m, _ := json.Marshal(maze)
	fmt.Println(maze)
	fmt.Println(m)
	w.Header().Set("Content-Type", "application/json")
	w.Write(m)
}

func main() {
	router := router.New()
	router.GET("/btree", btreeMaze)

	fmt.Println("listen localhost:3333")
	// http.HandleFunc("/", handler)
	http.ListenAndServe(":3333", router)

}
