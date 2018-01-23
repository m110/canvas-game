package board_test

import (
	"testing"

	"github.com/m110/canvas-game/pkg/game/board"
	"github.com/m110/canvas-game/pkg/game/events"
	"github.com/stretchr/testify/assert"
)

type StubPublisher struct{}

func (StubPublisher) Publish(event events.Event) {}

func TestNewBoard(t *testing.T) {
	b := board.NewBoard(StubPublisher{})
	player := b.Players()
	assert.Empty(t, player)
}

func TestSpawnPlayer(t *testing.T) {
	b := board.NewBoard(StubPublisher{})
	b.SpawnPlayer("")
}
