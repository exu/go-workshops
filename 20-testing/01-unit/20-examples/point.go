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
