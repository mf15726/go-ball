package main

import (
	"log"

	"image/color"

	"go.uber.org/dig"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mf15726/go-ball/assets"
	"github.com/mf15726/go-ball/player"
)

type Game struct{}

const SCREEN_WIDTH = 1280
const SCREEN_HEIGHT = 768

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	assets.DrawCircle(screen, SCREEN_WIDTH/2, SCREEN_HEIGHT/2, 50, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func main() {
	container := dig.New()
	container.Provide(player.PlayerRepositoryInMemoryImpl)
	container.Provide(player.PlayerServiceImpl)
	ebiten.SetWindowSize(1280, 768)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
