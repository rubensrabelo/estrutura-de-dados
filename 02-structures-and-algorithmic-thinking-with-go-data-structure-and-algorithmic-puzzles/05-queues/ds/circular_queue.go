package ds

import "errors"

var (
	ErrQueueEmpty = errors.New("ds: queue is empty")
	ErrQueueFull  = errors.New("ds: queue is full")
)

type CircularArrayQueue struct {
	front    int
	rear     int
	size     int
	capacity int
	array    []interface{}
}

func NewCircularArrayQueue(capacity int) *CircularArrayQueue {
	return &CircularArrayQueue{
		front:    0,
		rear:     -1,
		size:     0,
		capacity: capacity,
		array:    make([]interface{}, capacity),
	}
}

func (q *CircularArrayQueue) EnQueue(data interface{}) (int, error) {
	if q.IsFull() {
		return 0, ErrQueueFull
	}
	q.rear = (q.rear + 1) % q.capacity
	q.array[q.rear] = data
	q.size++
	return q.size, nil
}

func (q *CircularArrayQueue) DeQueue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, ErrQueueEmpty
	}
	removed := q.array[q.front]
	q.array[q.front] = nil
	q.front = (q.front + 1) % q.capacity
	q.size--
	return removed, nil
}

func (q *CircularArrayQueue) Front() (interface{}, error) {
	if q.IsEmpty() {
		return nil, ErrQueueEmpty
	}
	return q.array[q.front], nil
}

func (q *CircularArrayQueue) Rear() (interface{}, error) {
	if q.IsEmpty() {
		return nil, ErrQueueEmpty
	}
	return q.array[q.rear], nil
}

func (q *CircularArrayQueue) Size() int {
	return q.size
}

func (q *CircularArrayQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *CircularArrayQueue) IsFull() bool {
	return q.size == q.capacity
}

func (q *CircularArrayQueue) ToSlice() []interface{} {
	res := make([]interface{}, q.size)
	for i := 0; i < q.size; i++ {
		res[i] = q.array[(q.front+i)%q.capacity]
	}
	return res
}
