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

type BinaryTree struct {
	Root *TreeNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{Root: nil}
}

func (t *BinaryTree) Insert(value int) int {
	t.Root = t.insertRecursive(t.Root, value)
	return t.Size(t.Root)
}

func (t *BinaryTree) insertRecursive(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return &TreeNode{Value: value}
	}
	if t.Height(node.Left) <= t.Height(node.Right) {
		node.Left = t.insertRecursive(node.Left, value)
	} else {
		node.Right = t.insertRecursive(node.Right, value)
	}
	return node
}

func (t *BinaryTree) Delete(value int) (int, error) {
	if t.Root == nil {
		return 0, ErrEmptyTree
	}
	var err error
	t.Root, err = t.deleteRecursive(t.Root, value)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (t *BinaryTree) deleteRecursive(node *TreeNode, value int) (*TreeNode, error) {
	if node == nil {
		return nil, ErrValueNotFound
	}
	if node.Value == value {
		if node.Left == nil && node.Right == nil {
			return nil, nil
		}
		if node.Left != nil {
			maxNode := t.findMax(node.Left)
			node.Value = maxNode.Value
			var err error
			node.Left, err = t.deleteRecursive(node.Left, maxNode.Value)
			return node, err
		} else {
			minNode := t.findMin(node.Right)
			node.Value = minNode.Value
			var err error
			node.Right, err = t.deleteRecursive(node.Right, minNode.Value)
			return node, err
		}
	}
	var err error
	node.Left, err = t.deleteRecursive(node.Left, value)
	if err == nil {
		return node, nil
	}
	node.Right, err = t.deleteRecursive(node.Right, value)
	return node, err
}

func (t *BinaryTree) findMax(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	max := node
	left := t.findMax(node.Left)
	right := t.findMax(node.Right)
	if left != nil && left.Value > max.Value {
		max = left
	}
	if right != nil && right.Value > max.Value {
		max = right
	}
	return max
}

func (t *BinaryTree) findMin(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	min := node
	left := t.findMin(node.Left)
	right := t.findMin(node.Right)
	if left != nil && left.Value < min.Value {
		min = left
	}
	if right != nil && right.Value < min.Value {
		min = right
	}
	return min
}

func (t *BinaryTree) Search(value int) (*TreeNode, error) {
	return t.searchRecursive(t.Root, value)
}

func (t *BinaryTree) searchRecursive(node *TreeNode, value int) (*TreeNode, error) {
	if node == nil {
		return nil, ErrValueNotFound
	}
	if node.Value == value {
		return node, nil
	}
	if res, err := t.searchRecursive(node.Left, value); err == nil {
		return res, nil
	}
	return t.searchRecursive(node.Right, value)
}

func (t *BinaryTree) TraverseInOrder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	t.TraverseInOrder(node.Left, res)
	*res = append(*res, node.Value)
	t.TraverseInOrder(node.Right, res)
}

func (t *BinaryTree) Size(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + t.Size(node.Left) + t.Size(node.Right)
}

func (t *BinaryTree) Height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftHeight := t.Height(node.Left)
	rightHeight := t.Height(node.Right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func (t *BinaryTree) FindMaxSumLevel() (int, int) {
	h := t.Height(t.Root)
	maxSum := -1000000
	maxLevel := 0
	for i := 1; i <= h; i++ {
		sum := t.getLevelSum(t.Root, i)
		if sum > maxSum {
			maxSum = sum
			maxLevel = i
		}
	}
	return maxLevel, maxSum
}

func (t *BinaryTree) getLevelSum(node *TreeNode, level int) int {
	if node == nil {
		return 0
	}
	if level == 1 {
		return node.Value
	}
	return t.getLevelSum(node.Left, level-1) + t.getLevelSum(node.Right, level-1)
}

func (t *BinaryTree) LCA(node, n1, n2 *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	if node == n1 || node == n2 {
		return node
	}
	leftLCA := t.LCA(node.Left, n1, n2)
	rightLCA := t.LCA(node.Right, n1, n2)
	if leftLCA != nil && rightLCA != nil {
		return node
	}
	if leftLCA != nil {
		return leftLCA
	}
	return rightLCA
}
