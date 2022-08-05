package main

import (
	"SiDCo/canvas"
)

func main() {
	cnv := canvas.NewCanvas(canvas.Point{X: 1, Y: 1}, canvas.Point{X: 10, Y: 10})
	cnv.DrawBoxWithTitle("Title")
}
