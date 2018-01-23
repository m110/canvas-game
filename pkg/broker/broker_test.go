package broker_test

import (
	"testing"
	"time"

	"github.com/m110/canvas-game/pkg/broker"
	"github.com/m110/canvas-game/pkg/game/player"
	"github.com/m110/canvas-game/pkg/game/position"
)

func TestBroker(t *testing.T) {
	b := broker.NewBroker()

	playerChan := b.Subscribe(player.PlayerEvent{})

	wait := make(chan struct{})

	go func() {
		for {
			select {
			case <-playerChan:
				wait <- struct{}{}
				return
			case <-time.After(time.Second * 1):
				t.Fatal("Timed out")
				wait <- struct{}{}
				return
			}
		}
	}()

	p := player.NewPlayer("player-id", position.NewPosition(0, 0))
	playerJoined := player.NewJoinedEvent(*p)
	b.Publish(playerJoined)

	<-wait
}
