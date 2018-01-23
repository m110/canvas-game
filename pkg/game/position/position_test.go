package position_test

import (
	"testing"

	"github.com/m110/canvas-game/pkg/game/position"
	"github.com/stretchr/testify/assert"
)

func TestNewPosition(t *testing.T) {
	p := position.NewPosition(10, 20)

	assert.Equal(t, p.X(), 10)
	assert.Equal(t, p.Y(), 20)
}

func TestNewPositionWithOffset(t *testing.T) {
	cases := []struct {
		Name     string
		OffsetX  int
		OffsetY  int
		Expected position.Position
	}{
		{
			"no_changes",
			0,
			0,
			position.NewPosition(10, 20),
		},
		{
			"offset_x",
			5,
			0,
			position.NewPosition(15, 20),
		},
		{
			"offset_y",
			0,
			-5,
			position.NewPosition(10, 15),
		},
		{
			"offset_x_and_y",
			5,
			-5,
			position.NewPosition(15, 15),
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			p := position.NewPosition(10, 20)
			newPosition := position.NewPositionWithOffset(p, c.OffsetX, c.OffsetY)
			assert.Equal(t, c.Expected, newPosition)
		})
	}
}
