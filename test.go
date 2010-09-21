package main

import (
	"fmt"
        "./geom"
)

func main() {

	p1 := &geom.Point2D{10,0}
	p2 := &geom.Point2D{20,10}
	p3 := &geom.Point2D{10,20}
	p4 := &geom.Point2D{0,10}
	p5 := &geom.Point2D{10,0}

  	pa := []geom.Point2D{*p1, *p2, *p3, *p4, *p5}

	ps1 := &geom.Polyline2D{pa};

	fmt.Println(ps1.Length())

	return
}
