// It's a package!
package point

import (
	"strconv"
)

// THE point struct
type Point struct {
	X int
	Y int
}

// Give me string representation od THE point
func (point *Point) String() string {
	return strconv.Itoa(point.X) + ":" + strconv.Itoa(point.Y)
}

// Point can be moved by another point
func (point *Point) Move(to Point) {
	point.X += to.X
	point.Y += to.Y
}

// Pokazać godoc dla tego package'a
// lokalnie u mnie z wystartowanym serwerem godoc na porcie 7777
// przez `godoc -http=":7799" -goroot=$GOPATH/ &`
// http://localhost:7799/pkg/github.com/exu/go-workshops/302-testing-unit-examples/
