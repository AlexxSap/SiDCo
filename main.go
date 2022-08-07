package main

import (
	"SiDCo/canvas"
	"time"
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

	for i := 0; i < 5; i++ {
		cnv1.DrawSample(i%2 == 0)
		cnv2.DrawSample(i%2 != 0)

		time.Sleep(1 * time.Second)

	}
}
