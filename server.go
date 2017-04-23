package main

import (
	"./cors"
	"./middleware"
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
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(maze); err != nil {
		panic(err)
	}
}

func main() {
	router := router.New()
	options := &cors.Options{
		AllowOrigins: "*",
		AllowMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders: []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding"},
	}
	corsMid := cors.New(options)

	bindedRouter := middleware.New(corsMid.Handler).Bind(router)

	router.POST("/btree", btreeMaze)
	fmt.Println("listen localhost:3333")
	http.ListenAndServe(":3333", bindedRouter)

}
