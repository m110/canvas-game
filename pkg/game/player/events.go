package player

import (
	"time"
)

func NewJoinedEvent(player Player) PlayerEvent {
	return PlayerEvent{"PlayerJoined", time.Now(), player}
}

func NewMovedEvent(player Player) PlayerEvent {
	return PlayerEvent{"PlayerMoved", time.Now(), player}
}

type PlayerEvent struct {
	name       string
	occurredOn time.Time
	player     Player
}

func (e PlayerEvent) Name() string {
	return e.name
}

func (e PlayerEvent) OccurredOn() time.Time {
	return e.occurredOn
}

func (e PlayerEvent) Player() Player {
	return e.player
}
