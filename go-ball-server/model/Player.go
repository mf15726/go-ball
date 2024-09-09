package model

type Player struct {
	Id   int
	Name string
	X    float64
	Y    float64
}

func MakePlayer(name string) Player {
	return Player{Name: name, X: 0.0, Y: 0.0}
}

func NewPlayer(name string) *Player {
	return &Player{Name: name, X: 0.0, Y: 0.0}
}
