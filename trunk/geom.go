/* =====================
      David Fernandes
   daugfernandes@aim.com
   ===================== */

package geom

import (
	"fmt"
	"math"
)

// Point2D
type Point2D struct {
	X float64
	Y float64
}

// Calc distance between 2 points
func (p1 *Point2D) Distance(p2 *Point2D) float64 {
	return math.Sqrt(math.Pow(p1.X-p2.X, 2) + math.Pow(p1.Y-p2.Y, 2))
}

func (p *Point2D) Centroid() *Point2D {
	return &Point2D{p.X,p.Y}
}

// Rotate point
func (p *Point2D) Rotate(center *Point2D, angle float64) *Point2D {
	d := p.Distance(center)
	return &Point2D{center.X + d*math.Cos(angle),
		center.Y + d*math.Sin(angle)}
}

func (p *Point2D) String() string {
	return fmt.Sprintf("Point2d->%s%f%s%f", "X:", p.X, ",Y:", p.Y)
}

/*===========================================================================
  Polyline2D
  ===========================================================================*/

type Polyline2D struct {
	Vertexes []Point2D
}

func (pl *Polyline2D) String() string {
	return fmt.Sprintf("Polyline2D->%d%s",len(pl.Vertexes)," points")
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

// Calc Centroid of polyline
func (pl *Polyline2D) Centroid() *Point2D {
	var tx, ty, tl, l float64 = 0, 0, 0, 0
	var pPrev Point2D
	for i, vtx := range pl.Vertexes {
		if i == 0 {
			pPrev = vtx
		} else {
			l = vtx.Distance(&pPrev)
			tl += l 
			// tx,ty (coords of mid-point) weighted by length
			tx += (vtx.X + pPrev.X) / 2 * l
			ty += (vtx.Y + pPrev.Y) / 2 * l
			pPrev = vtx
		}
	}
	return &Point2D{tx / tl, ty / tl}
}
