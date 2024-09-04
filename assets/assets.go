package assets

import (
	"embed"
	"image"
	"image/color"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed *
var assets embed.FS

// var PlayerSprite = mustLoadImage("player.png")

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

// DrawCircle draws a circle on the given image at (cx, cy) with the given radius.
func DrawCircle(img *ebiten.Image, cx, cy, radius float64, clr color.Color) {
	for y := -radius; y <= radius; y++ {
		for x := -radius; x <= radius; x++ {
			if x*x+y*y <= radius*radius {
				img.Set(int(cx+x), int(cy+y), clr)
			}
		}
	}
}
