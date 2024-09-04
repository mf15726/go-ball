package main

import (
	"go.uber.org/dig"

	"github.com/mf15726/go-ball/player"
)

func main() {
	container := dig.New()
	container.Provide(player.PlayerRepositoryInMemoryImpl)
	container.Provide(player.PlayerServiceImpl)
}
