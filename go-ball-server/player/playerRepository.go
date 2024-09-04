package player

type PlayerRepository interface {
	GetByID(id int) (*Player, error)
	GetAll() []Player
	Create(player Player) (*Player, error)
	Update(player Player) error
	Delete(id int) error
}

type PlayerRepositoryInMemory struct {
	playerMap map[int]Player
	nextId    int
}

func (playerRepository *PlayerRepositoryInMemory) GetByID(id int) (*Player, error) {
	player := playerRepository.playerMap[id]
	return &player, nil
}

func (playerRepository *PlayerRepositoryInMemory) GetAll() []Player {
	values := make([]Player, 0, len(playerRepository.playerMap))

	for _, value := range values {
		values = append(values, value)
	}
	return values
}

func (playerRepository *PlayerRepositoryInMemory) Create(player Player) (*Player, error) {
	player.id = playerRepository.nextId
	playerRepository.nextId += 1
	playerRepository.playerMap[player.id] = player
	return &player, nil
}

func (playerRepository *PlayerRepositoryInMemory) Update(player Player) error {
	playerRepository.playerMap[player.id] = player
	return nil
}

func (playerRepository *PlayerRepositoryInMemory) Delete(id int) error {
	delete(playerRepository.playerMap, id)
	return nil
}

func PlayerRepositoryInMemoryImpl() PlayerRepository {
	playerRepository := PlayerRepositoryInMemory{make(map[int]Player), 1}
	return &playerRepository
}
