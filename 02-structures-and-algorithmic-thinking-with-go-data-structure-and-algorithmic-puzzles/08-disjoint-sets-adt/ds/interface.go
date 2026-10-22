package ds

import "errors"

var ErrElementNotFound = errors.New("element not found in the disjoint set")

type DisjointSet interface {
	MakeSet(value int)
	Find(value int) (int, error)
	Union(value1, value2 int) (bool, error)
	Size() int
	CountSets() int
}
