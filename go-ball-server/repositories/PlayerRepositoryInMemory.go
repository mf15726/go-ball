package repositories

import (
	"github.com/mf15726/go-ball/interfaces"
	"github.com/mf15726/go-ball/model"
)

type PlayerRepositoryInMemory struct {
	playerMap map[int]model.Player
	nextId    int
}

func (playerRepository *PlayerRepositoryInMemory) GetByID(id int) (*model.Player, error) {
	player := playerRepository.playerMap[id]
	return &player, nil
}

func (playerRepository *PlayerRepositoryInMemory) GetAll() []model.Player {
	values := make([]model.Player, 0, len(playerRepository.playerMap))

	for _, value := range values {
		values = append(values, value)
	}
	return values
}

func (playerRepository *PlayerRepositoryInMemory) Create(player model.Player) (*model.Player, error) {
	player.Id = playerRepository.nextId
	playerRepository.nextId += 1
	playerRepository.playerMap[player.Id] = player
	return &player, nil
}

func (playerRepository *PlayerRepositoryInMemory) Update(player model.Player) error {
	playerRepository.playerMap[player.Id] = player
	return nil
}

func (playerRepository *PlayerRepositoryInMemory) Delete(id int) error {
	delete(playerRepository.playerMap, id)
	return nil
}

func PlayerRepositoryInMemoryImpl() interfaces.IPlayerRepository {
	return &PlayerRepositoryInMemory{make(map[int]model.Player), 1}
}
