/*
 *
 *     Copyright (C) 2010  David Fernandes
 *
 *                         Rua da Quinta Amarela, 60
 *                         4475-663 MAIA
 *                         PORTUGAL
 *
 *                         <daugfernandes@aim.com>
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

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

/* Calculate distance between 2 points
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
	return fmt.Sprintf("Point2d->%s%f% s%f", "X:", this.X, ",Y:", this.Y)
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
     	return Length(this.Vertexes)
}

func Length(arr []Point2D) float64 {

     if len(arr) <  2 { return 0 }
     if len(arr) == 2 { return arr[0].Distance(&arr[1]) } 
     if len(arr) == 3 { return Length(arr[0:2])+Length(arr[1:3]) }
     // TODO: chanel
     return Length(arr[0:len(arr)/2])+Length(arr[len(arr)/2-1:len(arr)])

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

/* Calc MBR (Minimum Bounding Rectangle
*/
func (this *Polyline2D) MBR() *Polygon2D {

	minx, miny := this.Vertexes[0].X, this.Vertexes[0].Y 
	maxx, maxy := minx, miny

	for _, vtx := range this.Vertexes[1:] {
		if vtx.X < minx { 
			minx = vtx.X 
		}
		if vtx.X > maxx { 
			maxx = vtx.X 
		}
		if vtx.Y < miny { 
			miny = vtx.Y 
		}
		if vtx.Y > maxy { 
			maxy = vtx.Y 
		}
	}

	return &Polygon2D{[]Point2D{Point2D{minx,miny},Point2D{maxx,miny},Point2D{maxx,maxy},Point2D{minx,maxy},Point2D{minx,miny}}};
}

/* TODO: Calc ConvexHull
*/
func (this *Polyline2D) ConvexHul() *Polygon2D {
	return nil
}

/*===========================================================================
  Polygon2D
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

// Calc length of polygon border
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
   If Area < 0 the vertexes are ordered clockwise
   Some systems consider this case as a HOLE and so should this package
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

/* Calc MBR (Minimum Bounding Rectangle)
*/
func (this *Polygon2D) MBR() *Polygon2D {

	// Ommit [0] as == [size-1]
	minx, miny := this.Vertexes[1].X, this.Vertexes[1].Y 
	maxx, maxy := minx, miny

	for _, vtx := range this.Vertexes[2:] {
		if vtx.X < minx { 
			minx = vtx.X 
		}
		if vtx.X > maxx { 
			maxx = vtx.X 
		}
		if vtx.Y < miny { 
			miny = vtx.Y 
		}
		if vtx.Y > maxy { 
			maxy = vtx.Y 
		}
	}

	return &Polygon2D{[]Point2D{Point2D{minx,miny},Point2D{maxx,miny},Point2D{maxx,maxy},Point2D{minx,maxy},Point2D{minx,miny}}};
}

/* TODO: Calc ConvexHull
*/
func (this *Polygon2D) ConvexHul() *Polygon2D {
	return nil
}
