package services

import (
	"github.com/mf15726/go-ball/interfaces"
	"github.com/mf15726/go-ball/model"
)

type PlayerService struct {
	repo interfaces.IPlayerRepository
}

func (service *PlayerService) GetPlayer(id int) (*model.Player, error) {
	return service.repo.GetByID(id)
}

func (service *PlayerService) CreatePlayer(name string) (*model.Player, error) {
	return service.repo.Create(model.MakePlayer(name))
}

func PlayerServiceImpl(repo interfaces.IPlayerRepository) interfaces.IPlayerService {
	return &PlayerService{repo}
}
