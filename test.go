package main

import (
	"fmt"
        "./geom"
)

func main() {

	p1 := &geom.Point2D{10,10}
	p2 := &geom.Point2D{10,20}
	p3 := &geom.Point2D{20,20}
	p4 := &geom.Point2D{20,10}
	
	ps1 := &geom.Polygon2D{[]geom.Point2D{*p1, *p2, *p3, *p4, *p1}};

	fmt.Println(ps1.Area())


	return
}
