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
	RawBoard     [][]Color
	Turns        int
	CurrentColor Color
}

func NewBoard() *Board {
	return &Board{RawBoard: make([][]Color, Size+2)}
}

func (b *Board) Init() {
}

func (b *Board) Move() {
}

func (b *Board) Undo() {
}

func (b *Board) Pass() {
}

func (b *Board) CheckMobility(disc Disc) Direction {
	if b.RawBoard[disc.x][disc.y] != Empty {
		return None
	}

	var x, y int
	var dir Direction

	if b.RawBoard[disc.x][disc.y-1] == -disc.color {
		x = disc.x
		y = disc.y - 2
		for b.RawBoard[x][y] == -disc.color {
			y--
		}
		if b.RawBoard[x][y] == disc.color {
			dir |= Upper
		}
	}

	if b.RawBoard[disc.x][disc.y+1] == -disc.color {
		x = disc.x
		y = disc.y + 2
		for b.RawBoard[x][y] == -disc.color {
			y++
		}
		if b.RawBoard[x][y] == disc.color {
			dir |= Lower
		}
	}

	if b.RawBoard[disc.x-1][disc.y] == -disc.color {
		x = disc.x - 2
		y = disc.y
		for b.RawBoard[x][y] == -disc.color {
			x--
		}
		if b.RawBoard[x][y] == disc.color {
			dir |= Left
		}
	}

	if b.RawBoard[disc.x+1][disc.y] == -disc.color {
		x = disc.x + 2
		y = disc.y
		for b.RawBoard[x][y] == -disc.color {
			x++
		}
		if b.RawBoard[x][y] == disc.color {
			dir |= Right
		}
	}

	if b.RawBoard[disc.x+1][disc.y-1] == -disc.color {
		x = disc.x + 2
		y = disc.y - 2
		for b.RawBoard[x][y] == -disc.color {
			x++
			y--
		}
		if b.RawBoard[x][y] == disc.color {
			dir |= UpperRight
		}
	}

	if b.RawBoard[disc.x-1][disc.y-1] == -disc.color {
		x = disc.x - 2
		y = disc.y - 2
		for b.RawBoard[x][y] == -disc.color {
			x--
			y--
		}
		if b.RawBoard[x][y] == disc.color {
			dir |= UpperLeft
		}
	}

	if b.RawBoard[disc.x-1][disc.y+1] == -disc.color {
		x = disc.x - 2
		y = disc.y + 2
		for b.RawBoard[x][y] == -disc.color {
			x--
			y++
		}
		if b.RawBoard[x][y] == disc.color {
			dir |= LowerLeft
		}
	}

	if b.RawBoard[disc.x+1][disc.y+1] == -disc.color {
		x = disc.x + 2
		y = disc.y + 2
		for b.RawBoard[x][y] == -disc.color {
			x++
			y++
		}
		if b.RawBoard[x][y] == disc.color {
			dir |= LowerRight
		}
	}

	return dir
}
