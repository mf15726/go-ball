package player

type Player struct {
	id   int
	Name string
	X    float64
	Y    float64
}

func makePlayer(name string, x float64, y float64) Player {
	return Player{Name: name, X: x, Y: y}
}
