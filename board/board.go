package board

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) *Point {
	return &Point{x, y}
}

func PointIs(cord string) *Point {
	x := cord[0] - 'a' + 1
	y := cord[1] - '1' + 1
	return NewPoint(int(x), int(y))
}

type Disc struct {
	Point
	color Color
}

func NewDisc(x, y int, color Color) *Disc {
	return &Disc{*NewPoint(x, y), color}
}

type ColorStorage struct {
	data []int
}

func (s *ColorStorage) Get(color Color) int {
	return s.data[color+1]
}
func (s *ColorStorage) Set(color Color, value int) {
	s.data[color+1] = value
}

type Board struct {
	RawBoard     [][]Color
	MovableDir   [][][]Direction
	MovablePos   [][]Disc
	Turns        int
	CurrentColor Color
	Discs        ColorStorage
	UpdateLog    [][]Disc
}

func NewBoard() *Board {
	board := Board{
		RawBoard:   make([][]Color, Size+2),
		MovablePos: make([][]Disc, MaxTurns+1),
		MovableDir: make([][][]Direction, MaxTurns+1),
		Discs:      ColorStorage{data: make([]int, 3)},
	}
	for i := 0; i < Size+2; i++ {
		board.RawBoard[i] = make([]Color, Size+2)
	}

	for i := 0; i < MaxTurns+1; i++ {
		board.MovablePos[i] = make([]Disc, 1)
		board.MovableDir[i] = make([][]Direction, Size+2)
		for j := 0; j < Size+2; j++ {
			board.MovableDir[i][j] = make([]Direction, Size+2)
		}
	}

	board.Init()
	return &board
}

func (b *Board) Init() {
	for x := 1; x <= Size; x++ {
		for y := 1; y <= Size; y++ {
			b.RawBoard[x][y] = Empty
		}
	}

	for y := 0; y < Size+2; y++ {
		b.RawBoard[0][y] = Wall
		b.RawBoard[Size+1][y] = Wall
	}

	for x := 0; x < Size+2; x++ {
		b.RawBoard[x][0] = Wall
		b.RawBoard[x][Size+1] = Wall
	}

	b.RawBoard[4][4] = White
	b.RawBoard[5][5] = White
	b.RawBoard[4][5] = Black
	b.RawBoard[5][4] = Black

	b.Discs.Set(Black, 2)
	b.Discs.Set(White, 2)
	b.Discs.Set(Empty, Size*Size-4)

	b.Turns = 0
	b.CurrentColor = Black

	b.UpdateLog = make([][]Disc, 0, 0)

	b.InitMovable()
}

func (b *Board) CountDisc(color Color) int {
	return b.Discs.Get(color)
}

func (b *Board) GetColor(p Point) Color {
	return b.RawBoard[p.x][p.y]
}

func (b *Board) GetMovablePos() []Disc {
	return b.MovablePos[b.Turns]
}

func (b *Board) GetUpdate() []Disc {
	if len(b.UpdateLog) == 0 {
		return make([]Disc, 0, 0)
	}
	return b.UpdateLog[len(b.UpdateLog)-1]
}

func (b *Board) Move(p Point) bool {
	if (p.x < 0 || p.x > Size) || (p.y < 0 || p.y > Size) || b.MovableDir[b.Turns][p.x][p.y] == None {
		return false
	}

	b.FlipDiscs(p)

	b.Turns++
	b.CurrentColor = -b.CurrentColor

	b.InitMovable()

	return true
}

func (b *Board) Undo() bool {
	if b.Turns == 0 {
		return false
	}

	b.CurrentColor = -b.CurrentColor
	update := b.popUpdateLog()

	if len(update) == 0 {
		b.MovablePos[b.Turns] = make([]Disc, 0, 0)
		for x := 1; x <= Size; x++ {
			for y := 1; y <= Size; y++ {
				b.MovableDir[b.Turns][x][y] = None
			}
		}
	} else {
		b.Turns--
		p := update[0]
		b.RawBoard[p.x][p.y] = Empty
		for i := 1; i < len(update); i++ {
			p = update[i]
			b.RawBoard[p.x][p.y] = -b.CurrentColor
		}

		discDiff := len(update)
		b.Discs.Set(b.CurrentColor, b.Discs.Get(b.CurrentColor)-discDiff)
		b.Discs.Set(-b.CurrentColor, b.Discs.Get(-b.CurrentColor)+(discDiff-1))
		b.Discs.Set(Empty, b.Discs.Get(Empty)+1)
	}
	return true
}

func (b *Board) popUpdateLog() []Disc {
	last := len(b.UpdateLog) - 1
	discs := b.UpdateLog[last]
	b.UpdateLog = b.UpdateLog[:last]
	return discs
}

func (b *Board) Pass() bool {
	if len(b.MovablePos) != 0 || b.IsGameOver() {
		return false
	}

	b.CurrentColor = -b.CurrentColor
	b.UpdateLog = append(b.UpdateLog, make([]Disc, 0, 0))

	b.InitMovable()
	return true
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

func (b *Board) FlipDiscs(p Point) {
	dir := b.MovableDir[b.Turns][p.x][p.y]
	update := make([]Disc, 0, 0)

	b.RawBoard[p.x][p.y] = b.CurrentColor
	update = append(update, *NewDisc(p.x, p.y, b.CurrentColor))

	if dir&Upper != None {
		for y := p.y - 1; b.CurrentColor != b.RawBoard[p.x][y]; y -= 1 {
			b.RawBoard[p.x][y] = b.CurrentColor
			update = append(update, *NewDisc(p.x, y, b.CurrentColor))
		}
	}

	if dir&Lower != None {
		for y := p.y + 1; b.CurrentColor != b.RawBoard[p.x][y]; y += 1 {
			b.RawBoard[p.x][y] = b.CurrentColor
			update = append(update, *NewDisc(p.x, y, b.CurrentColor))
		}
	}

	if dir&Left != None {
		for x := p.x - 1; b.CurrentColor != b.RawBoard[x][p.y]; x -= 1 {
			b.RawBoard[x][p.y] = b.CurrentColor
			update = append(update, *NewDisc(x, p.y, b.CurrentColor))
		}
	}

	if dir&Right != None {
		for x := p.x + 1; b.CurrentColor != b.RawBoard[x][p.y]; x += 1 {
			b.RawBoard[x][p.y] = b.CurrentColor
			update = append(update, *NewDisc(x, p.y, b.CurrentColor))
		}
	}

	if dir&UpperRight != None {
		x := p.x + 1
		y := p.y - 1
		for b.CurrentColor != b.RawBoard[x][y] {
			b.RawBoard[x][y] = b.CurrentColor
			update = append(update, *NewDisc(x, y, b.CurrentColor))
			x += 1
			y -= 1
		}
	}

	if dir&UpperLeft != None {
		x := p.x - 1
		y := p.y - 1
		for b.CurrentColor != b.RawBoard[x][y] {
			b.RawBoard[x][y] = b.CurrentColor
			update = append(update, *NewDisc(x, y, b.CurrentColor))
			x -= 1
			y -= 1
		}
	}

	if dir&LowerLeft != None {
		x := p.x - 1
		y := p.y + 1
		for b.CurrentColor != b.RawBoard[x][y] {
			b.RawBoard[x][y] = b.CurrentColor
			update = append(update, *NewDisc(x, y, b.CurrentColor))
			x -= 1
			y += 1
		}
	}

	if dir&LowerRight != None {
		x := p.x + 1
		y := p.y + 1
		for b.CurrentColor != b.RawBoard[x][y] {
			b.RawBoard[x][y] = b.CurrentColor
			update = append(update, *NewDisc(x, y, b.CurrentColor))
			x += 1
			y += 1
		}
	}

	discDiff := len(update)
	b.Discs.Set(b.CurrentColor, b.Discs.Get(b.CurrentColor)+discDiff)
	b.Discs.Set(-b.CurrentColor, b.Discs.Get(-b.CurrentColor)-(discDiff-1))
	b.Discs.Set(Empty, b.Discs.Get(Empty)-1)

	b.UpdateLog = append(b.UpdateLog, update)
}

func (b *Board) InitMovable() {

	var dir Direction

	b.MovablePos[b.Turns] = make([]Disc, 0, 0)

	for x := 1; x <= Size; x++ {
		disc := NewDisc(0, 0, b.CurrentColor)
		disc.x = x
		for y := 1; y <= Size; y++ {
			disc.y = y
			dir = b.CheckMobility(*disc)
			if dir != None {
				b.MovablePos[b.Turns] = append(b.MovablePos[b.Turns], *disc)
			}
			b.MovableDir[b.Turns][x][y] = dir
		}
	}
}

func (b *Board) IsGameOver() bool {
	if b.Turns == MaxTurns {
		return true
	}
	if len(b.MovablePos[b.Turns]) != 0 {
		return false
	}

	disc := NewDisc(0, 0, -b.CurrentColor)
	for x := 0; x < Size; x++ {
		disc.x = x
		for y := 0; y < Size; y++ {
			disc.y = y
			if b.CheckMobility(*disc) != None {
				return false
			}
		}
	}
	return true
}
