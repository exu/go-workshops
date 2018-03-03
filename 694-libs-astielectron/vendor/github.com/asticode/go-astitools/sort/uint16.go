package astisort

import "sort"

// Uint16 sorts a slice of uint16s in increasing order.
func Uint16(a []uint16) { sort.Sort(UInt16Slice(a)) }

// UInt16Slice attaches the methods of Interface to []uint16, sorting in increasing order.
type UInt16Slice []uint16

func (p UInt16Slice) Len() int           { return len(p) }
func (p UInt16Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p UInt16Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
