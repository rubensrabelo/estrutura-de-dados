package ds

type TraversalService struct {
	tree BinaryTree
}

func NewTraversalService(tree BinaryTree) *TraversalService {
	return &TraversalService{tree: tree}
}

func (s *TraversalService) PreOrder() []int {
	var res []int
	s.preOrderRecursive(s.tree.GetRoot(), &res)
	return res
}

func (s *TraversalService) preOrderRecursive(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	*res = append(*res, node.Value)
	s.preOrderRecursive(node.Left, res)
	s.preOrderRecursive(node.Right, res)
}

func (s *TraversalService) InOrder() []int {
	var res []int
	s.inOrderRecursive(s.tree.GetRoot(), &res)
	return res
}

func (s *TraversalService) inOrderRecursive(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	s.inOrderRecursive(node.Left, res)
	*res = append(*res, node.Value)
	s.inOrderRecursive(node.Right, res)
}

func (s *TraversalService) PostOrder() []int {
	var res []int
	s.postOrderRecursive(s.tree.GetRoot(), &res)
	return res
}

func (s *TraversalService) postOrderRecursive(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	s.postOrderRecursive(node.Left, res)
	s.postOrderRecursive(node.Right, res)
	*res = append(*res, node.Value)
}

func (s *TraversalService) LevelOrder() []int {
	root := s.tree.GetRoot()
	if root == nil {
		return nil
	}

	var res []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		res = append(res, node.Value)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return res
}
