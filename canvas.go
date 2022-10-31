package SiDCo

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

var (
	// max line for move cursor after draw
	maxLine int = 0
)

type Color int

const (
	ColorReset Color = iota
	ColorBlack
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
)

func (color Color) String() string {
	return []string{
		"\u001b[0m",
		"\u001b[30m",
		"\u001b[31m",
		"\u001b[32m",
		"\u001b[33m",
		"\u001b[34m",
		"\u001b[35m",
		"\u001b[36m",
		"\u001b[37m"}[color]
}

// constants for comands
const (
	hLine             = "\u2500"
	vLine             = "\u2502"
	movePattern       = "\033[%v;%vf"
	topLeftCorner     = "\u250C"
	topRightCorner    = "\u2510"
	bottomLeftCorner  = "\u2514"
	bottomRightCorner = "\u2518"
	clearAllScreen    = "\u001b[2J"
)

// saveMaxLine save maximum line number into 'maxLine'
func saveMaxLine(line int) {
	if maxLine < line {
		maxLine = line
	}
}

// clearScreen clear all terminal screen
func clearScreen() {
	fmt.Println(clearAllScreen)
}

func init() {
	clearScreen()
}

// Point represents position in console
type Point struct {
	Line, Column int
}

// Canvas is the canvas for drawing
type Canvas struct {
	start Point
	size  Point
}

// NewCanvas create canvas with custom start point and size
func NewCanvas(start, size Point) (Canvas, error) {

	if size.Column == 0 || size.Line == 0 {
		return Canvas{}, fmt.Errorf("invalid size")
	}

	saveMaxLine(start.Line + size.Line)
	return Canvas{
		start: start,
		size:  size}, nil
}

// Size resurn size of box
func (cnv Canvas) Size() Point {
	return cnv.size
}

// PointProvider interface for provide point and symbol for draw
type PointProvider interface {
	Provide() (string, int, int)
}

// DrawSource draw symbol on position provided by PointProvider
func (cnv Canvas) DrawSource(pp PointProvider) {
	sym, line, col := pp.Provide()
	cnv.moveCursorTo(Point{line, col})
	fmt.Print(sym)
}

// EndDraw invoke when the drawing is end and need to put cursor pos to the bottom of screen
func (cnv Canvas) EndDraw() {
	cnv.moveCursorTo(Point{Line: maxLine + 1, Column: 1})
}

// DrawPath draw symbol 'sym' in point 'points'
func (cnv Canvas) DrawPath(sym string, points []Point) {
	for _, p := range points {
		cnv.moveCursorTo(p)
		fmt.Print(sym)
	}

	cnv.moveCursorTo(Point{Line: maxLine + 1, Column: 1})
}

// DrawText print text on specified position
func (cnv Canvas) DrawText(text string, point Point) {
	cnv.DrawColoredText(text, point, ColorReset)
}

// DrawColoredText print colored text on specified position
func (cnv Canvas) DrawColoredText(text string, point Point, color Color) {
	line, column := point.Line, point.Column

	fmt.Print(color.String())
	defer fmt.Print(ColorReset.String())

	for ind, sym := range []rune(text) {
		cnv.moveCursorTo(Point{Line: line, Column: column + ind})
		fmt.Print(string(sym))
	}

	cnv.moveCursorTo(Point{Line: maxLine + 1, Column: 1})
}

// ClearInner clear all in the box
func (cnv Canvas) ClearInner() {
	for line := 1; line < cnv.size.Line; line++ {
		cnv.moveCursorTo(Point{Line: line, Column: 1})
		fmt.Print(strings.Repeat(" ", int(cnv.size.Column)))
	}
	cnv.moveCursorTo(Point{Line: maxLine + 1, Column: 1})
}

// moveCursorTo moved cursor to custom position
func (cnv Canvas) moveCursorTo(point Point) {
	fmt.Printf(movePattern, cnv.start.Line+point.Line, cnv.start.Column+point.Column)
}

func (cnv Canvas) moveCursorToBottom() {
	fmt.Printf(movePattern, maxLine+2, 1)
}

// DrawBoxWithTitle draw box around canvas with some title
func (cnv Canvas) DrawBoxWithTitle(title string) error {
	return cnv.DrawColoredBoxWithTitle(title, ColorReset, ColorReset)
}

// DrawColoredBoxWithTitle draw box with spicified color around canvas with some colored title
func (cnv Canvas) DrawColoredBoxWithTitle(title string, boxColor Color, textColor Color) error {

	titleLen := utf8.RuneCountInString(title)
	if titleLen > int(cnv.size.Column) {
		return fmt.Errorf("length of title more then box width")
	}

	defer fmt.Print(ColorReset.String())

	cnv.moveCursorTo(Point{0, 0})

	// draw top
	fmt.Print(boxColor.String())
	fmt.Print(topLeftCorner + strings.Repeat(hLine, 2))

	fmt.Print(textColor.String())
	fmt.Print(title)

	fmt.Print(boxColor.String())
	fmt.Print(strings.Repeat(hLine, int(cnv.size.Column)-2-utf8.RuneCountInString(title)) +
		topRightCorner)

	for line := 1; line < cnv.size.Line; line++ {
		cnv.moveCursorTo(Point{Line: line, Column: 0})
		fmt.Print(vLine)
		cnv.moveCursorTo(Point{Line: line, Column: cnv.size.Column + 1})
		fmt.Print(vLine)
	}

	// draw bottom
	cnv.moveCursorTo(Point{Line: cnv.size.Line, Column: 0})
	fmt.Print(
		bottomLeftCorner +
			strings.Repeat(hLine, int(cnv.size.Column)) +
			bottomRightCorner)

	cnv.moveCursorToBottom()

	return nil

}
