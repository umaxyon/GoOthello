package main

import (
	"Osero/board"
	"Osero/console"
	"fmt"
)

func main() {
	b := board.NewBoard()

	for true {
		console.Print(*b)
		fmt.Println()
		fmt.Printf("Black %d  White %d  Empty %d",
			b.CountDisc(board.Black), b.CountDisc(board.White), b.CountDisc(board.Empty))
		fmt.Println()

		if b.IsGameOver() {
			fmt.Println("--------Game Over!!---------")
			return
		}

		var in string
		fmt.Print("Input (p: pass, u: undo): ")
		fmt.Scan(&in)

		if in == "p" {
			if !b.Pass() {
				fmt.Print("\nCan't Pass!!\n\n")
			}
			continue
		}

		if in == "u" {
			b.Undo()
			continue
		}

		p := board.PointIs(in)

		if !b.Move(*p) {
			fmt.Print("\nCan't put there!!\n\n")
			continue
		}
	}
}
