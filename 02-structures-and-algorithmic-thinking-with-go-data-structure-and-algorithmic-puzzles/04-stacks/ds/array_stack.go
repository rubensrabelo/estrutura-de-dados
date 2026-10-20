package ds

import "errors"

var (
	ErrStackOverflow  = errors.New("ds: stack overflow")
	ErrStackUnderflow = errors.New("ds: stack underflow")
)

type Stack struct {
	top      int
	size     int
	capacity uint
	array    []interface{}
}

func NewStack(capacity uint) *Stack {
	return &Stack{
		top:      -1,
		size:     0,
		capacity: capacity,
		array:    make([]interface{}, capacity),
	}
}

func (s *Stack) Push(data interface{}) (int, error) {
	if s.IsFull() {
		return s.size, ErrStackOverflow
	}
	s.top++
	s.array[s.top] = data
	s.size++
	return s.size, nil
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrStackUnderflow
	}
	removed := s.array[s.top]
	s.array[s.top] = nil
	s.top--
	s.size--
	return removed, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrStackUnderflow
	}
	return s.array[s.top], nil
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) IsFull() bool {
	return s.size == int(s.capacity)
}

func (s *Stack) ToSlice() []interface{} {
	return s.array[:s.size]
}
