package interfaces

import "github.com/mf15726/go-ball/model"

type IPlayerService interface {
	GetPlayer(int) (*model.Player, error)
	CreatePlayer(string) (*model.Player, error)
}
