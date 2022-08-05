package canvas

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

var (
	maxLine int = 0
)

func saveMaxLine(line int) {
	if maxLine < line {
		maxLine = line
	}
}

// constants for comands
const (
	hLine = "\u2500"
	vLine = "\u2502"
	//erazer            = "\033[H\033[2J"
	movePattern       = "\033[%v;%vf"
	topLeftCorner     = "\u250C"
	topRightCorner    = "\u2510"
	bottomLeftCorner  = "\u2514"
	bottomRightCorner = "\u2518"

//	clearAllScreen    = "\u001b[2J"
)

// func ClearScreen() {
// 	fmt.Println(clearAllScreen)
// }

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

	saveMaxLine(start.Line + size.Line)
	return Canvas{
		start: start,
		size:  size}
}

// clear clear canvas
func (cnv Canvas) clear() {
	/// TODO fill me
}

// clearInner clear all in the box
// func (cnv Canvas) clearInner() {
// 	/// TODO fill me
// }

// clearFullCurentLine will clear screen on current line
// func (cnv Canvas) clearFullCurentLine() {
// 	fmt.Println(erazer)
// }

// moveCursorTo moved cursor to custom position
func (cnv Canvas) moveCursorTo(point Point) {
	fmt.Printf(movePattern, point.Line, point.Column)
}

// DrawBoxWithTitle draw box around canvas with some title
func (cnv Canvas) DrawBoxWithTitle(title string) {

	/// TODO add check to title's len

	cnv.clear()

	cnv.moveCursorTo(cnv.start)

	// draw top
	fmt.Print(
		topLeftCorner +
			strings.Repeat(hLine, 2) +
			title +
			strings.Repeat(hLine, int(cnv.size.Column-2-utf8.RuneCountInString(title))) +
			topRightCorner)

	for line := cnv.start.Line + 1; line < cnv.start.Line+cnv.size.Line; line++ {
		cnv.moveCursorTo(Point{Line: line, Column: cnv.start.Column})
		fmt.Print(vLine)
		cnv.moveCursorTo(Point{Line: line, Column: cnv.start.Column + cnv.size.Column + 1})
		fmt.Print(vLine)
	}

	// draw bottom
	cnv.moveCursorTo(Point{Line: cnv.start.Line + cnv.size.Line, Column: cnv.start.Column})
	fmt.Print(
		bottomLeftCorner +
			strings.Repeat(hLine, cnv.size.Column) +
			bottomRightCorner)

	cnv.moveCursorTo(Point{Line: maxLine + 1, Column: 1})

}
