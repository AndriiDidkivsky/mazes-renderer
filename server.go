package main

import (
	"./router"
	"encoding/json"
	"fmt"
	"github.com/AndriiDidkivsky/go-algorithms/btree_maze"
	"net/http"
)

type MazeBody struct {
	Width, Height int
}

func btreeMaze(w http.ResponseWriter, r *http.Request) {
	//write statuses; refactoring
	var body MazeBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		fmt.Println(err)
	}
	maze := btree_maze.GenerateMaze(body.Width, body.Height)
	m, _ := json.Marshal(maze)
	w.Header().Set("Content-Type", "application/json")
	w.Write(m)
}

func main() {
	router := router.New()
	router.POST("/btree", btreeMaze)

	fmt.Println("listen localhost:3333")
	// http.HandleFunc("/", handler)
	http.ListenAndServe(":3333", router)

}
