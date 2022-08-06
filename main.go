package main

import (
	"SiDCo/canvas"
)

func main() {

	cnv1, err := canvas.NewCanvas(canvas.Point{Line: 1, Column: 1}, canvas.Point{Line: 10, Column: 20})
	if err != nil {
		panic(err)
	}
	cnv1.DrawBoxWithTitle("Title1")

	cnv2, err := canvas.NewCanvas(canvas.Point{Line: 1, Column: 25}, canvas.Point{Line: 5, Column: 30})
	if err != nil {
		panic(err)
	}
	cnv2.DrawBoxWithTitle("Title2")
}
