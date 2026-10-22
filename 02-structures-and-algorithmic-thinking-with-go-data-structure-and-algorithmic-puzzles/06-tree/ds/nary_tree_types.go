package ds

import "errors"

var (
	ErrNaryValueNotFound = errors.New("ds: n-ary value not found")
	ErrNaryEmptyTree     = errors.New("ds: n-ary tree is empty")
)

type NaryTreeNode struct {
	Value    int
	Children []*NaryTreeNode
}

type NaryTree interface {
	GetRoot() *NaryTreeNode
	Insert(parentValue, newValue int) (int, error)
	Delete(value int) (int, error)
	Search(value int) (*NaryTreeNode, error)
	Size() int
	Height() int
	FindMaxSumLevel() (int, int)
	LCA(n1, n2 *NaryTreeNode) *NaryTreeNode
}
