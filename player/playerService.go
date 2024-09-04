package player

type PlayerService struct {
	repo PlayerRepository
}

func (service *PlayerService) GetPlayer(id int) (*Player, error) {
	return service.repo.GetByID(id)
}

func (service *PlayerService) CreatePlayer(name string) (*Player, error) {
	return service.repo.Create(makePlayer(name, 0, 0))
}

func PlayerServiceImpl(repo PlayerRepository) *PlayerService {
	return &PlayerService{repo: repo}
}
