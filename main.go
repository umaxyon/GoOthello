package main

import (
	"Osero/board"
	"Osero/console"
)

func main() {
	b := board.NewBoard()
	console.NewConsole(*b).Start()
}
