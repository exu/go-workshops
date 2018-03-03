package astisort

import "sort"

// Uint8 sorts a slice of uint8s in increasing order.
func Uint8(a []uint8) { sort.Sort(UInt8Slice(a)) }

// UInt8Slice attaches the methods of Interface to []uint8, sorting in increasing order.
type UInt8Slice []uint8

func (p UInt8Slice) Len() int           { return len(p) }
func (p UInt8Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p UInt8Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
