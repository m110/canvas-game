package broker_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/m110/canvas-game/pkg/broker"
	"github.com/m110/canvas-game/pkg/game/player"
)

func TestBroker(t *testing.T) {
	b := broker.NewBroker()

	joinedChan := b.Subscribe(player.PlayerEvent{})
	movedChan := b.Subscribe(player.MovedEvent{})

	wait := make(chan struct{})

	go func() {
		for {
			select {
			case joinedEvent := <-joinedChan:
				fmt.Println("Joined", joinedEvent)
			case movedEvent := <-movedChan:
				fmt.Println("Moved", movedEvent)
			case <-time.After(time.Second * 1):
				fmt.Println("timed out")
				wait <- struct{}{}
				return
			}
		}
	}()

	playerJoined := player.NewJoinedEvent("player-id")
	playerMoved := player.NewMovedEvent("player-id", player.NewPosition(1, 2))

	b.Publish(playerJoined)
	b.Publish(playerMoved)

	<-wait
}
