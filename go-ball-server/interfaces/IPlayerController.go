package interfaces

import (
	"net/http"

	"github.com/mf15726/go-ball/model"
)

type IPlayerController interface {
	CreatePlayer(http.ResponseWriter, *http.Request) (*model.Player, error)
	HelloWorld(http.ResponseWriter, *http.Request)
}
