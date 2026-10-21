package ds

import "errors"

var (
	ErrValueNotFound = errors.New("ds: value not found")
	ErrEmptyTree     = errors.New("ds: tree is empty")
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTree interface {
	GetRoot() *TreeNode
	Insert(value int) int
	Delete(value int) (int, error)
	Search(value int) (*TreeNode, error)
	Size() int
	Height() int
	FindMaxSumLevel() (int, int)
	LCA(n1, n2 *TreeNode) *TreeNode
}
