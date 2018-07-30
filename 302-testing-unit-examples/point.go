// Package point represents points in 2-dimensional environment
package point

import (
	"strconv"
)

// Point - THE point represents point in two dimensional coordinates
type Point struct {
	X int
	Y int
}

// String - Gives me string representation of THE point
func (point *Point) String() string {
	return strconv.Itoa(point.X) + ":" + strconv.Itoa(point.Y)
}

// Move - Point can be moved by another point
func (point *Point) Move(to Point) {
	point.X += to.X
	point.Y += to.Y
}

// show godoc for this package
// locally with godoc on port 7777
// przez `godoc -http=":7799" -goroot=$GOPATH/ &`
// http://localhost:7799/pkg/github.com/exu/go-workshops/302-testing-unit-examples/
