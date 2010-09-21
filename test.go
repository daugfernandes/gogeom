package main

import (
	"fmt"
        "./geom"
)

func main() {

	p1 := &geom.Point2D{10,10}
	p2 := &geom.Point2D{20,20}
	p3 := &geom.Point2D{30,10}

  	pa := []geom.Point2D{*p1, *p2, *p3}

	ps1 := &geom.Polyline2D{pa};

	fmt.Println(ps1.Centroid().String())
	//fmt.Println(ps1.Length(),ps1.String())
	//fmt.Println("++",p1.Centroid().String())

	return
}
