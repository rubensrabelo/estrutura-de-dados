package ds

type stackNode struct {
	data interface{}
	next *stackNode
}

type LinkedListStack struct {
	top  *stackNode
	size int
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{
		top:  nil,
		size: 0,
	}
}

func (s *LinkedListStack) Push(data interface{}) (int, error) {
	newNode := &stackNode{
		data: data,
		next: s.top,
	}
	s.top = newNode
	s.size++
	return s.size, nil
}

func (s *LinkedListStack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrStackUnderflow
	}
	removed := s.top.data
	s.top = s.top.next
	s.size--
	return removed, nil
}

func (s *LinkedListStack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrStackUnderflow
	}
	return s.top.data, nil
}

func (s *LinkedListStack) Size() int {
	return s.size
}

func (s *LinkedListStack) IsEmpty() bool {
	return s.size == 0
}

func (s *LinkedListStack) ToSlice() []interface{} {
	res := make([]interface{}, s.size)
	curr := s.top
	for i := 0; curr != nil; i++ {
		res[i] = curr.data
		curr = curr.next
	}
	return res
}
