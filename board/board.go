package board

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) *Point {
	return &Point{x, y}
}

type Disc struct {
	Point
	color Color
}

func NewDisc(x, y int, color Color) *Disc {
	return &Disc{*NewPoint(x, y), color}
}

type Board struct {
	RawBoard     [][]int
	Turns        int
	CurrentColor Color
}

func NewBoard() *Board {
	return &Board{RawBoard: make([][]int, Size+2)}
}

func (b *Board) Init() {
}

func (b *Board) Move() {
}

func (b *Board) Undo() {
}

func (b *Board) Pass() {
}
