package events

import "time"

type Event interface {
	Name() string
	OccurredOn() time.Time
}
