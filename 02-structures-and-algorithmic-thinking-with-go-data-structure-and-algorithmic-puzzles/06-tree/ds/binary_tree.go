package ds

type SimpleBinaryTree struct {
	Root *TreeNode
}

func NewSimpleBinaryTree() *SimpleBinaryTree {
	return &SimpleBinaryTree{Root: nil}
}

func (t *SimpleBinaryTree) GetRoot() *TreeNode {
	return t.Root
}

func (t *SimpleBinaryTree) Insert(value int) int {
	t.Root = t.insertRecursive(t.Root, value)
	return t.Size()
}

func (t *SimpleBinaryTree) insertRecursive(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return &TreeNode{Value: value}
	}
	if t.heightRecursive(node.Left) <= t.heightRecursive(node.Right) {
		node.Left = t.insertRecursive(node.Left, value)
	} else {
		node.Right = t.insertRecursive(node.Right, value)
	}
	return node
}

func (t *SimpleBinaryTree) Delete(value int) (int, error) {
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

func (t *SimpleBinaryTree) deleteRecursive(node *TreeNode, value int) (*TreeNode, error) {
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

func (t *SimpleBinaryTree) findMax(node *TreeNode) *TreeNode {
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

func (t *SimpleBinaryTree) findMin(node *TreeNode) *TreeNode {
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

func (t *SimpleBinaryTree) Search(value int) (*TreeNode, error) {
	return t.searchRecursive(t.Root, value)
}

func (t *SimpleBinaryTree) searchRecursive(node *TreeNode, value int) (*TreeNode, error) {
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

func (t *SimpleBinaryTree) Size() int {
	return t.sizeRecursive(t.Root)
}

func (t *SimpleBinaryTree) sizeRecursive(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + t.sizeRecursive(node.Left) + t.sizeRecursive(node.Right)
}

func (t *SimpleBinaryTree) Height() int {
	return t.heightRecursive(t.Root)
}

func (t *SimpleBinaryTree) heightRecursive(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftHeight := t.heightRecursive(node.Left)
	rightHeight := t.heightRecursive(node.Right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func (t *SimpleBinaryTree) FindMaxSumLevel() (int, int) {
	h := t.Height()
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

func (t *SimpleBinaryTree) getLevelSum(node *TreeNode, level int) int {
	if node == nil {
		return 0
	}
	if level == 1 {
		return node.Value
	}
	return t.getLevelSum(node.Left, level-1) + t.getLevelSum(node.Right, level-1)
}

func (t *SimpleBinaryTree) LCA(n1, n2 *TreeNode) *TreeNode {
	return t.lcaRecursive(t.Root, n1, n2)
}

func (t *SimpleBinaryTree) lcaRecursive(node, n1, n2 *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	if node == n1 || node == n2 {
		return node
	}
	leftLCA := t.lcaRecursive(node.Left, n1, n2)
	rightLCA := t.lcaRecursive(node.Right, n1, n2)
	if leftLCA != nil && rightLCA != nil {
		return node
	}
	if leftLCA != nil {
		return leftLCA
	}
	return rightLCA
}
