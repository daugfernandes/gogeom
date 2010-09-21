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

/* Calc distance between 2 points
*/
func (this *Point2D) Distance(p *Point2D) float64 {
	return math.Sqrt(math.Pow(this.X-p.X, 2) + math.Pow(this.Y-p.Y, 2))
}

func (this *Point2D) Centroid() *Point2D {
	return &Point2D{this.X,this.Y}
}

// Rotate point
func (this *Point2D) Rotate(center *Point2D, angle float64) *Point2D {
	d := this.Distance(center)
	return &Point2D{center.X + d*math.Cos(angle),
		center.Y + d*math.Sin(angle)}
}

func (this *Point2D) String() string {
	return fmt.Sprintf("Point2d->%s%f%s%f", "X:", this.X, ",Y:", this.Y)
}

/* Compares two points
   TODO: Maybe some tolerance value should be considered!
*/
func (this *Point2D) Equal(p *Point2D) bool {
	return this.X == p.X && this.Y==p.Y
}

/*===========================================================================
  Polyline2D
  ===========================================================================*/

type Polyline2D struct {
	Vertexes []Point2D
}

func (this *Polyline2D) String() string {
	return fmt.Sprintf("Polyline2D->%d%s",this.Size()," points")
}

func (this *Polyline2D) Size() int {
	return len(this.Vertexes)
}

// Calc length of polyline
func (this *Polyline2D) Length() float64 {
	var result float64 = 0;
	var pPrev Point2D
	for i, vtx := range this.Vertexes {
		if i != 0 {
			fmt.Println(vtx.String(),pPrev.String())
			result += vtx.Distance(&pPrev)
		}
		pPrev = vtx
	}
	return result
}

// Calc Centroid of polyline
func (this *Polyline2D) Centroid() *Point2D {
	var tx, ty, tl, l float64 = 0, 0, 0, 0
	var pPrev Point2D
	for i, vtx := range this.Vertexes {
		if i != 0 {
			l = vtx.Distance(&pPrev)
			tl += l 
			// tx,ty (coords of mid-point weighted by length)
			tx += (vtx.X + pPrev.X) / 2 * l
			ty += (vtx.Y + pPrev.Y) / 2 * l
		}
		pPrev = vtx
	}
	return &Point2D{tx / tl, ty / tl}
}

/* Compares two polylines
   TODO: Maybe some tolerance value should be considered!
*/
func (this *Polyline2D) Equal(pl *Polyline2D) bool {
	if this.Size() != pl.Size() {
		return false
	}

	// Equal size
	for i, vtx := range this.Vertexes {
		if !vtx.Equal(&pl.Vertexes[i]) {
			return false
		}
	}
	return true
}

/*===========================================================================
  PolygonD
  TODO: Consider the possibility of "HOLES"
  ===========================================================================*/

type Polygon2D struct {
	Vertexes []Point2D
}

func (this *Polygon2D) String() string {
	return fmt.Sprintf("Polygon2D->%d%s",this.Size()," points")
}

func (this *Polygon2D) Size() int {
	return len(this.Vertexes)
}

// Calc length of polygon
func (this *Polygon2D) Length() float64 {
	var result float64 = 0;
	var pPrev Point2D
	for i, vtx := range this.Vertexes {
		if i != 0 {
			fmt.Println(vtx.String(),pPrev.String())
			result += vtx.Distance(&pPrev)
		}
		pPrev = vtx
	}
	return result
}

/* Calc Area of Polygon
   If Area < 0 the vertexes are ordered counter clockwise
   Some systems consider this case as a HOLE
*/
   
func (this *Polygon2D) Area() float64 {
	var t float64 = 0
	var pPrev Point2D
	for i, vtx := range this.Vertexes {
		if i != 0 {
			t += (pPrev.X * vtx.Y - vtx.X * pPrev.Y)
		}
		pPrev = vtx
	}
	return t / 2
}


// Calc Centroid of Polygon
func (this *Polygon2D) Centroid() *Point2D {
	var tx, ty, tl, l float64 = 0, 0, 0, 0
	var pPrev Point2D
	for i, vtx := range this.Vertexes {
		if i != 0 {
			l = vtx.Distance(&pPrev)
			tl += l 
			// tx,ty (coords of mid-point weighted by length)
			tx += (vtx.X + pPrev.X) / 2 * l
			ty += (vtx.Y + pPrev.Y) / 2 * l
		}
		pPrev = vtx
	}
	return &Point2D{tx / tl, ty / tl}
}

/* Compares two Polygon
   Somehow different from the Polyline case: the Vertexes array can be rotated:
   say (1,2,3,4,5,1) == (3,4,5,1,2,3) for this matter

   TODO: Maybe some tolerance value should be considered!
*/
func (this *Polygon2D) Equal(pl *Polygon2D) bool {
	if this.Size() != pl.Size() {
		return false
	}

	size := len(this.Vertexes) - 1
	i1, i2, consecutive_matches := 0, 0, 0

	for total_comparisons := 0; total_comparisons < 2 * size && consecutive_matches < size; total_comparisons ++ {

		if this.Vertexes[i1].Equal(&pl.Vertexes[i2]) {
			consecutive_matches ++
			i1++
			if i1>=size {
				i1=0
			}
		} else {
			consecutive_matches = 0
		}

		i2 ++
		if i2>=size {
			i2=0
		}
	}

	return consecutive_matches == size
}


