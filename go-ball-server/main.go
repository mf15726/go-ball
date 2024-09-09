package main

import (
	"log"
	"net/http"

	"github.com/mf15726/go-ball/controllers"
	"github.com/mf15726/go-ball/interfaces"
	"github.com/mf15726/go-ball/repositories"
	"github.com/mf15726/go-ball/services"
)

func main() {
	var playerController interfaces.IPlayerController
	playerController, _, _ = createPlayerContainers()

	http.HandleFunc("/hello-world", playerController.HelloWorld)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createPlayerContainers() (interfaces.IPlayerController, interfaces.IPlayerService, interfaces.IPlayerRepository) {
	playerRepository := repositories.PlayerRepositoryInMemoryImpl()
	playerService := services.PlayerServiceImpl(playerRepository)
	playerController := controllers.PlayerControllerImpl(playerService)
	return playerController, playerService, playerRepository
}

func createPlayerEndpoints(playerController interfaces.IPlayerController) {
	http.HandleFunc("/hello-world", playerController.HelloWorld)
	// http.HandleFunc("/player/create", playerController.CreatePlayer)
}
