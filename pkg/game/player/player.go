package player

import (
	"github.com/m110/canvas-game/pkg/game/position"
)

type Player struct {
	id       string
	position position.Position
}

func NewPlayer(id string, position position.Position) *Player {
	return &Player{
		id:       id,
		position: position,
	}
}

func (p Player) ID() string {
	return p.id
}

func (p Player) Position() position.Position {
	return p.position
}

func (p *Player) Move(direction string) {
	modX := 0
	modY := 0

	switch direction {
	case "n":
		modY = -1
	case "e":
		modX = -1
	case "s":
		modY = 1
	case "w":
		modX = 1
	}

	p.position = position.NewPositionWithOffset(p.position, modX, modY)
}
