package controllers

import (
	"fmt"
	"net/http"

	"github.com/mf15726/go-ball/interfaces"
	"github.com/mf15726/go-ball/model"
)

type PlayerController struct {
	service interfaces.IPlayerService
}

func (controller *PlayerController) CreatePlayer(w http.ResponseWriter, r *http.Request) (*model.Player, error) {
	return controller.service.CreatePlayer("Test")
}

func (controller *PlayerController) HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func PlayerControllerImpl(service interfaces.IPlayerService) interfaces.IPlayerController {
	return &PlayerController{service}
}
