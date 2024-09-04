package main

import (
	"go.uber.org/dig"
	"github.com/mf15726/go-ball/player"

	"net/http"
)

func main() {
	createContainers()
	
	mux := http.NewServeMux()

	// Register the routes and handlers
	mux.Handle("/", &homeHandler{})

	// Run the server
	http.ListenAndServe(":8080", mux)
}

func createContainers() {
	container := dig.New()
	container.Provide(player.PlayerRepositoryInMemoryImpl)
	container.Provide(player.PlayerServiceImpl)
}

func 
