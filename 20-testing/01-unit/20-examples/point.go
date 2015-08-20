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
func (this *Point) String() string {
	return strconv.Itoa(this.X) + ":" + strconv.Itoa(this.Y)
}

// Point can be moved by another point
func (this *Point) Move(point Point) {
	this.X += point.X
	this.Y += point.Y
}

// Pokazać godoc dla tego package'a
// lokalnie u mnie z wystartowanym serwerem godoc na porcie 7777
// przez `godoc -http ':7777'&`
// http://l:7777/pkg/github.com/exu/go-workshops/20-testing/01-unit/20-examples/
