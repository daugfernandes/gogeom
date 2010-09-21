/* David Fernandes
   daugfernandes@aim.com */

package geom

import (
	"fmt"
	"math"
)

type Point2D struct {
	X float64
	Y float64
}

type Polyline2D struct {
	Vertexes []Point2D
}

// Calc length of polyline
func (pl *Polyline2D) Length() float64 {
	var result float64 = 0;
	var pPrev Point2D
	for i, vtx := range pl.Vertexes {
		if i == 0 {
			pPrev = vtx
		} else {
			fmt.Println(vtx.String(),pPrev.String())
			result += vtx.Distance(&pPrev)
			pPrev = vtx
		}
	}
	return result
}

// Calc distance between 2 points
func (p1 *Point2D) Distance(p2 *Point2D) float64 {
	return math.Sqrt(math.Pow(p1.X-p2.X, 2) + math.Pow(p1.Y-p2.Y, 2))
}

// Rotate point
func (p1 *Point2D) Rotate(center *Point2D, angle float64) *Point2D {
	d := p1.Distance(center)
	return &Point2D{center.X + d*math.Cos(angle),
		center.Y + d*math.Sin(angle)}
}

func (p1 *Point2D) String() string {
	return fmt.Sprintf("%s%f%s%f", "X:", p1.X, ",Y:", p1.Y)
}

func (pl *Polyline2D) String() string {
	return "x"
}
