package position

type Position struct {
	x int
	y int
}

func NewPosition(x, y int) Position {
	return Position{x, y}
}

func NewPositionWithOffset(p Position, x, y int) Position {
	return Position{
		x: p.X() + x,
		y: p.Y() + y,
	}
}

func (p Position) X() int {
	return p.x
}

func (p Position) Y() int {
	return p.y
}
