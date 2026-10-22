package ds

type SimpleNaryTree struct {
	Root *NaryTreeNode
}

func NewSimpleNaryTree() *SimpleNaryTree {
	return &SimpleNaryTree{Root: nil}
}

func (t *SimpleNaryTree) GetRoot() *NaryTreeNode {
	return t.Root
}

func (t *SimpleNaryTree) Insert(parentValue, newValue int) (int, error) {
	newNode := &NaryTreeNode{Value: newValue, Children: make([]*NaryTreeNode, 0)}
	if t.Root == nil {
		t.Root = newNode
		return t.Size(), nil
	}

	parent, err := t.Search(parentValue)
	if err != nil {
		return 0, err
	}

	parent.Children = append(parent.Children, newNode)
	return t.Size(), nil
}

func (t *SimpleNaryTree) Delete(value int) (int, error) {
	if t.Root == nil {
		return 0, ErrNaryEmptyTree
	}

	if t.Root.Value == value {
		t.Root = nil
		return value, nil
	}

	deleted, err := t.deleteRecursive(t.Root, value)
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

func (t *SimpleNaryTree) deleteRecursive(node *NaryTreeNode, value int) (int, error) {
	if node == nil {
		return 0, ErrNaryValueNotFound
	}

	for i, child := range node.Children {
		if child.Value == value {
			node.Children = append(node.Children[:i], node.Children[i+1:]...)
			node.Children = append(node.Children, child.Children...)
			return value, nil
		}
	}

	for _, child := range node.Children {
		if deleted, err := t.deleteRecursive(child, value); err == nil {
			return deleted, nil
		}
	}

	return 0, ErrNaryValueNotFound
}

func (t *SimpleNaryTree) Search(value int) (*NaryTreeNode, error) {
	return t.searchRecursive(t.Root, value)
}

func (t *SimpleNaryTree) searchRecursive(node *NaryTreeNode, value int) (*NaryTreeNode, error) {
	if node == nil {
		return nil, ErrNaryValueNotFound
	}
	if node.Value == value {
		return node, nil
	}
	for _, child := range node.Children {
		if res, err := t.searchRecursive(child, value); err == nil {
			return res, nil
		}
	}
	return nil, ErrNaryValueNotFound
}

func (t *SimpleNaryTree) Size() int {
	return t.sizeRecursive(t.Root)
}

func (t *SimpleNaryTree) sizeRecursive(node *NaryTreeNode) int {
	if node == nil {
		return 0
	}
	size := 1
	for _, child := range node.Children {
		size += t.sizeRecursive(child)
	}
	return size
}

func (t *SimpleNaryTree) Height() int {
	return t.heightRecursive(t.Root)
}

func (t *SimpleNaryTree) heightRecursive(node *NaryTreeNode) int {
	if node == nil {
		return 0
	}
	maxChildHeight := 0
	for _, child := range node.Children {
		h := t.heightRecursive(child)
		if h > maxChildHeight {
			maxChildHeight = h
		}
	}
	return maxChildHeight + 1
}

func (t *SimpleNaryTree) FindMaxSumLevel() (int, int) {
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

func (t *SimpleNaryTree) getLevelSum(node *NaryTreeNode, level int) int {
	if node == nil {
		return 0
	}
	if level == 1 {
		return node.Value
	}
	sum := 0
	for _, child := range node.Children {
		sum += t.getLevelSum(child, level-1)
	}
	return sum
}

// LCA encontra o Menor Ancestral Comum (Lowest Common Ancestor) entre dois nós.
func (t *SimpleNaryTree) LCA(n1, n2 *NaryTreeNode) *NaryTreeNode {
	return t.lcaRecursive(t.Root, n1, n2)
}

func (t *SimpleNaryTree) lcaRecursive(node, n1, n2 *NaryTreeNode) *NaryTreeNode {
	if node == nil {
		return nil
	}
	if node == n1 || node == n2 {
		return node
	}

	count := 0
	var tempNode *NaryTreeNode

	for _, child := range node.Children {
		res := t.lcaRecursive(child, n1, n2)
		if res != nil {
			count++
			tempNode = res
		}
	}

	if count > 1 {
		return node
	}
	return tempNode
}
