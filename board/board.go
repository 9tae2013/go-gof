package board

import (
	"fmt"
	"math/rand"
)

type Board [][]Cell

func NewBoard(x, y int) Board {
	return genBoard(x, y)
}

func NewBoardGlider() Board {
	return [][]Cell{
		{DeadCell{X: 0, Y: 0}, DeadCell{X: 0, Y: 1}, DeadCell{X: 0, Y: 2}, DeadCell{X: 0, Y: 3}, DeadCell{X: 0, Y: 4}},
		{DeadCell{X: 1, Y: 0}, DeadCell{X: 1, Y: 1}, LiveCell{X: 1, Y: 2}, DeadCell{X: 1, Y: 3}, DeadCell{X: 1, Y: 4}},
		{DeadCell{X: 2, Y: 0}, DeadCell{X: 2, Y: 1}, DeadCell{X: 2, Y: 2}, LiveCell{X: 2, Y: 3}, DeadCell{X: 2, Y: 4}},
		{DeadCell{X: 3, Y: 0}, LiveCell{X: 3, Y: 1}, LiveCell{X: 3, Y: 2}, LiveCell{X: 3, Y: 3}, DeadCell{X: 3, Y: 4}},
		{DeadCell{X: 4, Y: 0}, DeadCell{X: 4, Y: 1}, DeadCell{X: 4, Y: 2}, DeadCell{X: 4, Y: 3}, DeadCell{X: 4, Y: 4}},
	}
}

func NewBoardBlinker() Board {
	return [][]Cell{
		{DeadCell{X: 0, Y: 0}, DeadCell{X: 0, Y: 1}, DeadCell{X: 0, Y: 2}, DeadCell{X: 0, Y: 3}, DeadCell{X: 0, Y: 4}},
		{DeadCell{X: 1, Y: 0}, DeadCell{X: 1, Y: 1}, DeadCell{X: 1, Y: 2}, DeadCell{X: 1, Y: 3}, DeadCell{X: 1, Y: 4}},
		{DeadCell{X: 2, Y: 0}, LiveCell{X: 2, Y: 1}, LiveCell{X: 2, Y: 2}, LiveCell{X: 2, Y: 3}, DeadCell{X: 2, Y: 4}},
		{DeadCell{X: 3, Y: 0}, DeadCell{X: 3, Y: 1}, DeadCell{X: 3, Y: 2}, DeadCell{X: 3, Y: 3}, DeadCell{X: 3, Y: 4}},
		{DeadCell{X: 4, Y: 0}, DeadCell{X: 4, Y: 1}, DeadCell{X: 4, Y: 2}, DeadCell{X: 4, Y: 3}, DeadCell{X: 4, Y: 4}},
	}
}

func (board Board) RefreshBoard() Board {
	var newboard Board
	for _i := range board {
		var row []Cell
		for _j := range board[_i] {
			c := board[_i][_j].NextStage(board)
			row = append(row, c)
		}
		newboard = append(newboard, row)
	}
	return newboard
}

func (board Board) Display() {
	fmt.Println("Size ", len(board), "x", len(board[0]))
	for _i := range board {
		for _j := range board[_i] {
			if board[_i][_j].Live() {
				fmt.Print("X")
			} else {
				fmt.Print(" ")
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func (board Board) NeighboursLive(c Cell) int {
	cells := neighbours(c, board)
	return countLive(cells)
}

func countLive(cells []Cell) int {
	l := 0
	for _, v := range cells {
		if v.Live() {
			l = l + 1
		}
	}
	return l
}

func neighbours(c Cell, board Board) []Cell {
	var cells []Cell
	_sx := len(board)
	_sy := len(board[0])
	for cx := -1; cx <= 1; cx = cx + 1 {
		for cj := -1; cj <= 1; cj = cj + 1 {
			_x := c.GetX() + cx
			_y := c.GetY() + cj
			if _x >= 0 && _x < _sx &&
				_y >= 0 && _y < _sy &&
				!(_x == c.GetX() && _y == c.GetY()) {
				cells = append(cells, board[_x][_y])
			}
		}
	}
	return cells
}

func genBoard(i int, j int) [][]Cell {
	cells := make([][]Cell, i)
	for i := range cells {
		cells[i] = make([]Cell, j)
	}
	return randomCell(cells)
}

func randomCell(cells [][]Cell) [][]Cell {
	for _i := range cells {
		for _j := range cells[_i] {
			if rand.Int()%2 == 0 {
				cells[_i][_j] = &LiveCell{
					X: _i,
					Y: _j,
				}
			} else {
				cells[_i][_j] = &DeadCell{
					X: _i,
					Y: _j,
				}
			}
		}
	}
	return cells
}
