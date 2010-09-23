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

package main

import (
	"fmt"
        "./geom"
)


func main() {
	
	p1 := &geom.Point2D{1.10,0.2}
	p2 := &geom.Point2D{2.033,-1.3120}
	p3 := &geom.Point2D{1.3032,-2.4067}
	p4 := &geom.Point2D{0,1.201}

	pl := []geom.Point2D{*p1,*p2,*p3,*p4}

	plg := &geom.Polygon2D{pl}

	fmt.Println(plg.Area())
	return
}
