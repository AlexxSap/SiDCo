package canvas

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// constants for comands
const (
	line              = "\u2500"
	erazer            = "\033[H\033[2J"
	movePattern       = "\033[%v;%vf"
	topLeftCorner     = "\u250C"
	topRightCorner    = "\u2510"
	bottomLeftCorner  = "\u2514"
	bottomRightCorner = "\u2518"
)

// Point represents position in console
type Point struct {
	/// TODO change to uint8 or 16
	Line, Column int
}

// Canvas is the canvas for drawing
type Canvas struct {
	start Point
	size  Point
}

// NewCanvas create canvas with custom start point and size
func NewCanvas(start, size Point) Canvas {
	/// TODO add size check
	return Canvas{
		start: start,
		size:  size}
}

// clear clear canvas
func (cnv Canvas) clear() {
	/// TODO fill me
}

// clearFullCurentLine will clear screen on current line
func (cnv Canvas) clearFullCurentLine() {
	fmt.Println(erazer)
}

// moveCursorTo moved cursor to custom position
func (cnv Canvas) moveCursorTo(point Point) {
	fmt.Printf(movePattern, point.Line, point.Column)
}

// DrawBoxWithTitle draw box around canvas with some title
func (cnv Canvas) DrawBoxWithTitle(title string) {
	/// TODO add check to title's len
	cnv.clear()

	/// TODO delete this
	cnv.clearFullCurentLine()
	cnv.moveCursorTo(cnv.start)

	// draw top
	fmt.Print(
		topLeftCorner +
			strings.Repeat(line, 2) +
			title +
			strings.Repeat(line, int(cnv.size.Line-2-utf8.RuneCountInString(title))) +
			topRightCorner)
	fmt.Println("")

	// draw bottom
	cnv.moveCursorTo(Point{Line: cnv.start.Line + cnv.size.Line, Column: cnv.start.Column})
	fmt.Print(
		bottomLeftCorner +
			strings.Repeat(line, cnv.size.Line) +
			bottomRightCorner)
	fmt.Println("")

}
