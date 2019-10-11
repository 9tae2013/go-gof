package board

type Cell interface {
	GetX() int
	GetY() int
	Live() bool
	NextStage(b Board) Cell
}

type DeadCell struct {
	X, Y int
}

func (c DeadCell) GetX() int {
	return c.X
}

func (c DeadCell) GetY() int {
	return c.Y
}

func (c DeadCell) Live() bool {
	return false
}

func (c DeadCell) NextStage(b Board) Cell {
	if b.NeighboursLive(c) == 3 {
		return &LiveCell{X: c.X, Y: c.Y}
	}
	return c
}

type LiveCell struct {
	X, Y int
}

func (c LiveCell) GetX() int {
	return c.X
}

func (c LiveCell) GetY() int {
	return c.Y
}

func (c LiveCell) Live() bool {
	return true
}

func (c LiveCell) NextStage(b Board) Cell {
	l := b.NeighboursLive(c)
	if l != 2 && l != 3 {
		return &DeadCell{X: c.X, Y: c.Y}
	}
	return c
}