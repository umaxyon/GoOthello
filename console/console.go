package console

import (
	"Osero/board"
	"fmt"
)

const (
	BlackDisc = "x"
	WhiteDisc = "o"
)

type Console struct {
	board board.Board
}

func NewConsole(b board.Board) *Console {
	return &Console{board: b}
}

func (c *Console) DrawCurrentDisc() string {
	var disc string
	switch c.board.CurrentColor {
	case board.Black:
		disc = BlackDisc
		break
	case board.White:
		disc = WhiteDisc
		break
	}
	return disc
}

func (c *Console) Print() {
	fmt.Println("  a b c d e f g h ")
	for y := 1; y <= 8; y++ {
		fmt.Printf("%d", y)
		for x := 1; x <= 8; x++ {
			switch c.board.GetColor(*board.NewPoint(x, y)) {
			case board.Black:
				fmt.Printf(" %s", BlackDisc)
				break
			case board.White:
				fmt.Printf(" %s", WhiteDisc)
				break
			default:
				fmt.Print("  ")
				break
			}
		}
		fmt.Println()
	}
}

func (c *Console) Start() {
	for true {
		b := &c.board
		c.Print()
		fmt.Println()
		fmt.Printf("Black %d  White %d  Empty %d",
			b.CountDisc(board.Black), b.CountDisc(board.White), b.CountDisc(board.Empty))
		fmt.Println()

		if b.IsGameOver() {
			fmt.Println("--------Game Over!!---------")
			return
		}

		var in string
		fmt.Printf("Input[%s] (p: pass, u: undo): ", c.DrawCurrentDisc())
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
