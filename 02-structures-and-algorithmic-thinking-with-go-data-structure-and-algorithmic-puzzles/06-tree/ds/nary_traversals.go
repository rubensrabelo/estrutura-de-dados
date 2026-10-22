package ds

type NaryTraversalService struct {
	tree NaryTree
}

func NewNaryTraversalService(tree NaryTree) *NaryTraversalService {
	return &NaryTraversalService{tree: tree}
}

func (s *NaryTraversalService) PreOrder() []int {
	var res []int
	s.preOrderRecursive(s.tree.GetRoot(), &res)
	return res
}

func (s *NaryTraversalService) preOrderRecursive(node *NaryTreeNode, res *[]int) {
	if node == nil {
		return
	}
	*res = append(*res, node.Value)
	for _, child := range node.Children {
		s.preOrderRecursive(child, res)
	}
}

func (s *NaryTraversalService) PostOrder() []int {
	var res []int
	s.postOrderRecursive(s.tree.GetRoot(), &res)
	return res
}

func (s *NaryTraversalService) postOrderRecursive(node *NaryTreeNode, res *[]int) {
	if node == nil {
		return
	}
	for _, child := range node.Children {
		s.postOrderRecursive(child, res)
	}
	*res = append(*res, node.Value)
}

func (s *NaryTraversalService) LevelOrder() []int {
	root := s.tree.GetRoot()
	if root == nil {
		return nil
	}

	var res []int
	queue := []*NaryTreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		res = append(res, node.Value)

		for _, child := range node.Children {
			if child != nil {
				queue = append(queue, child)
			}
		}
	}

	return res
}
