package board

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/m110/canvas-game/pkg/game/events"
	"github.com/m110/canvas-game/pkg/game/player"
	"github.com/m110/canvas-game/pkg/game/position"
)

type Publisher interface {
	Publish(event events.Event)
}

type Board struct {
	players map[string]*player.Player
	broker  Publisher
}

func NewBoard(broker Publisher) *Board {
	return &Board{
		players: map[string]*player.Player{},
		broker:  broker,
	}
}

func (b *Board) SpawnPlayer(playerID string) {
	_, ok := b.players[playerID]
	if ok {
		// TOOD return error
	}

	pos := position.NewPosition(rand.Intn(60), rand.Intn(40))
	p := player.NewPlayer(playerID, pos)

	b.players[playerID] = p
	b.broker.Publish(player.NewJoinedEvent(*p))
	fmt.Println("Added player", p)
}

func (b Board) Players() []player.Player {
	var all []player.Player
	for _, p := range b.players {
		all = append(all, *p)
	}
	return all
}

func (b Board) MovePlayer(playerID, direction string) {
	log.Println("Moving player", playerID, "in direction", direction)

	p, ok := b.players[playerID]
	if !ok {
		// TODO return error
	}

	p.Move(direction)
	b.broker.Publish(player.NewMovedEvent(*p))
}
