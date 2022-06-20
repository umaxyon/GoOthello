package console

import (
	"Osero/board"
	"fmt"
)

func Print(b board.Board) {
	fmt.Println("  a b c d e f g h ")
	for y := 1; y <= 8; y++ {
		fmt.Printf("%d", y)
		for x := 1; x <= 8; x++ {
			switch b.GetColor(*board.NewPoint(x, y)) {
			case board.Black:
				fmt.Print(" x")
				break
			case board.White:
				fmt.Print(" o")
				break
			default:
				fmt.Print("  ")
				break
			}
		}
		fmt.Println()
	}
}
