# SiDCo
SImple Drawing in Console 

### Installation
Install and update this package with `go get -u github.com/AlexxSap/SiDCo`

### Example
```go
package main

import (
	"time"
	canvas "github.com/AlexxSap/SiDCo"
)

func drawSample(cnv *canvas.Canvas, isHorizontal bool) {

	cnv.ClearInner()

	topLeft, bottomRight := cnv.TopLeftCorner(), cnv.BottomRightCorner()
	lineSize, columnSize := bottomRight.Line-topLeft.Line, bottomRight.Column-topLeft.Column

	if isHorizontal {
		points1, points2 := make([]canvas.Point, 0), make([]canvas.Point, 0)
		for i := 1; i <= columnSize; i++ {
			points1 = append(points1, canvas.Point{Line: topLeft.Line + 1, Column: topLeft.Column + i})
			points2 = append(points2, canvas.Point{Line: bottomRight.Line - 1, Column: topLeft.Column + i})
		}
		cnv.DrawPath("*", points1)
		cnv.DrawPath("+", points2)

	} else {

		points1, points2 := make([]canvas.Point, 0), make([]canvas.Point, 0)
		for i := 1; i < lineSize; i++ {
			points1 = append(points1, canvas.Point{Line: topLeft.Line + i, Column: topLeft.Column + 1})
			points2 = append(points2, canvas.Point{Line: topLeft.Line + i, Column: bottomRight.Column})
		}

		cnv.DrawPath("^", points1)
		cnv.DrawPath("#", points2)
	}
}

func main() {

	cnv1, err := canvas.NewCanvas(canvas.Point{Line: 1, Column: 1}, canvas.Point{Line: 10, Column: 20})
	if err != nil {
		panic(err)
	}
	cnv1.DrawBoxWithTitle("Title Square")

	cnv2, err := canvas.NewCanvas(canvas.Point{Line: 1, Column: 25}, canvas.Point{Line: 5, Column: 30})
	if err != nil {
		panic(err)
	}
	cnv2.DrawBoxWithTitle("Title Rectangle")

	for i := 0; i < 5; i++ {
		drawSample(&cnv1, i%2 == 0)
		drawSample(&cnv2, i%2 != 0)

		time.Sleep(1 * time.Second)
	}
}
```
![sample](https://user-images.githubusercontent.com/13485922/184956461-c1317a2f-e220-4116-94bb-9bdebbafda1e.gif)

