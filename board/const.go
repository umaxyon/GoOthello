package board

type Color int
type Direction int

const (
	Size int = 8

	Black Color = 1
	Empty Color = 0
	White Color = -1
	Wall  Color = 2

	None       Direction = 0
	Upper      Direction = 1
	UpperLeft  Direction = 2
	Left       Direction = 4
	LowerLeft  Direction = 8
	Lower      Direction = 16
	LowerRight Direction = 32
	Right      Direction = 64
	UpperRight Direction = 128
)
