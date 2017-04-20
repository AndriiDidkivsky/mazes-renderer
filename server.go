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
	fmt.Println(maze)
	// m, _ := json.Marshal(maze)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(maze); err != nil {
		panic(err)
	}
}

func main() {
	router := router.New()
	//think about cors middleware
	router.OPTIONS("/btree", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
	})
	router.POST("/btree", btreeMaze)

	fmt.Println("listen localhost:3333")
	// http.HandleFunc("/", handler)
	http.ListenAndServe(":3333", router)

}
