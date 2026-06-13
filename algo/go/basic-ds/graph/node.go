package graph

import (
	"github.com/bizshuk/algo/util"
)

type Node struct {
	Val       int
	X         int
	Y         int
	Neighbors []*Node
}

func CalDist(x1, y1, x2, y2 int) int {
	return util.AbsDiff(x1, x2) + util.AbsDiff(y1, y2)
}

func CalNodeDist(a, b *Node) int {
	return util.AbsDiff(a.X, a.X) + util.AbsDiff(b.Y, b.Y)
}
