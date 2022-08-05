package main

import (
	"SiDCo/canvas"
)

func main() {
	cnv := canvas.NewCanvas(canvas.Point{Line: 1, Column: 1}, canvas.Point{Line: 10, Column: 50})
	cnv.DrawBoxWithTitle("Title")
}
