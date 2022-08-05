package main

import (
	"SiDCo/canvas"
)

func main() {
	/// TODO clear screen before draw
	// canvas.ClearScreen()

	cnv1 := canvas.NewCanvas(canvas.Point{Line: 1, Column: 1}, canvas.Point{Line: 10, Column: 20})
	cnv1.DrawBoxWithTitle("Title1")

	cnv2 := canvas.NewCanvas(canvas.Point{Line: 1, Column: 25}, canvas.Point{Line: 5, Column: 30})
	cnv2.DrawBoxWithTitle("Title2")
}
