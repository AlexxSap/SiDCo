# SiDCo
[![Go Report Card](https://goreportcard.com/badge/github.com/AlexxSap/SiDCo)](https://goreportcard.com/report/github.com/AlexxSap/SiDCo)
[![Go Reference](https://pkg.go.dev/badge/github.com/AlexxSap/SiDCo.svg)](https://pkg.go.dev/github.com/AlexxSap/SiDCo)

SImple Drawing in Console. Not very much features, but everything you need sometimes.

See example below.

### Installation
Install and update this package with `go get -u github.com/AlexxSap/SiDCo`

### Example 1 (for v1.0.0)
```go
package main

import (
	"time"

	canvas "github.com/AlexxSap/SiDCo"
)

func drawSample(cnv *canvas.Canvas, isHorizontal bool) {

	cnv.ClearInner()

	lineSize, columnSize := cnv.Size().Line, cnv.Size().Column

	if isHorizontal {
		points1, points2 := make([]canvas.Point, 0), make([]canvas.Point, 0)
		for i := 1; i <= columnSize; i++ {
			points1 = append(points1, canvas.Point{1, i})
			points2 = append(points2, canvas.Point{lineSize - 1, i})
		}
		cnv.DrawPath("*", points1)
		cnv.DrawPath("+", points2)

	} else {

		points1, points2 := make([]canvas.Point, 0), make([]canvas.Point, 0)
		for i := 1; i < lineSize; i++ {
			points1 = append(points1, canvas.Point{i, 1})
			points2 = append(points2, canvas.Point{i, columnSize - 1})
		}

		cnv.DrawPath("^", points1)
		cnv.DrawPath("#", points2)
	}
}

func main() {

	cnv1, err := canvas.NewCanvas(canvas.Point{1, 1}, canvas.Point{10, 20})
	if err != nil {
		panic(err)
	}
	cnv1.DrawBoxWithTitle("Simple box")

	cnv2, err := canvas.NewCanvas(canvas.Point{1, 25}, canvas.Point{5, 30})
	if err != nil {
		panic(err)
	}
	cnv2.DrawColoredBoxWithTitle("Colored box", canvas.ColorRed, canvas.ColorGreen)

	for i := 0; i < 5; i++ {
		drawSample(&cnv1, i%2 == 0)
		drawSample(&cnv2, i%2 != 0)

		cnv2.DrawColoredText("constant text", canvas.Point{2, 3}, canvas.ColorYellow)

		time.Sleep(1 * time.Second)
	}
}


```
![sample2](https://user-images.githubusercontent.com/13485922/185736039-7098a5b0-d317-4bf7-9514-eb8ffd876d5e.gif)

### Example 2
Game snake - https://github.com/AlexxSap/snake


