package ds

type Node struct {
	Value  int
	Parent *Node
	Rank   int
}

type UnionFind struct {
	elements map[int]*Node
}

func NewUnionFind() *UnionFind {
	return &UnionFind{
		elements: make(map[int]*Node),
	}
}

func (uf *UnionFind) MakeSet(value int) {
	if _, exists := uf.elements[value]; exists {
		return
	}
	node := &Node{
		Value: value,
		Rank:  0,
	}
	node.Parent = node
	uf.elements[value] = node
}

func (uf *UnionFind) Find(value int) (int, error) {
	node, exists := uf.elements[value]
	if !exists {
		return 0, ErrElementNotFound
	}
	root := uf.findRepresentative(node)
	return root.Value, nil
}

func (uf *UnionFind) findRepresentative(node *Node) *Node {
	if node.Parent == node {
		return node
	}

	node.Parent = uf.findRepresentative(node.Parent)
	return node.Parent
}

func (uf *UnionFind) Union(value1, value2 int) (bool, error) {
	node1, exists1 := uf.elements[value1]
	node2, exists2 := uf.elements[value2]

	if !exists1 || !exists2 {
		return false, ErrElementNotFound
	}

	root1 := uf.findRepresentative(node1)
	root2 := uf.findRepresentative(node2)

	if root1 == root2 {
		return false, nil
	}

	if root1.Rank < root2.Rank {
		root1.Parent = root2
	} else if root1.Rank > root2.Rank {
		root2.Parent = root1
	} else {
		root2.Parent = root1
		root1.Rank++
	}

	return true, nil
}

func (uf *UnionFind) Size() int {
	return len(uf.elements)
}

func (uf *UnionFind) CountSets() int {
	count := 0
	for _, node := range uf.elements {
		if node.Parent == node {
			count++
		}
	}
	return count
}
