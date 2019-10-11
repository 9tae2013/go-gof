package main

import (
	"github.com/buger/goterm"
	"gof/board"
	"time"
)

func main() {
	var b = board.NewBoard(20, 20)
	goterm.Clear()
	goterm.Flush()
	time.Sleep(time.Second * 1)
	for {
		goterm.Clear()
		goterm.Flush()

		b = b.RefreshBoard()
		b.Display()
		time.Sleep(time.Second * 1)
	}
}
