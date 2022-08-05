package canvas

// constants for comands
const (
	line        = "\u2500"
	erazer      = "\033[H\033[2J"
	movePattern = "\033[%v;%vf"
)

// Point represents position in console
type Point struct {
	X, Y uint8
}

// Canvas is the canvas for drawing
type Canvas struct {
	start Point
	size  Point
}

// NewCanvas create canvas with custom start point and size
func NewCanvas(start, size Point) Canvas {
	return Canvas{
		start: start,
		size:  size}
}

func (cnv Canvas) DrawBoxWithTitle(title string) {

}
