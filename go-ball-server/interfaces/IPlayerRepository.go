package interfaces

import "github.com/mf15726/go-ball/model"

type IPlayerRepository interface {
	GetByID(id int) (*model.Player, error)
	GetAll() []model.Player
	Create(player model.Player) (*model.Player, error)
	Update(player model.Player) error
	Delete(id int) error
}
