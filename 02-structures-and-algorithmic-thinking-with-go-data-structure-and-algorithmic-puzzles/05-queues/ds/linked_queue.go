package ds

type queueNode struct {
	data interface{}
	next *queueNode
}

type LinkedListQueue struct {
	head *queueNode
	tail *queueNode
	size int
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (q *LinkedListQueue) EnQueue(data interface{}) (int, error) {
	newNode := &queueNode{data: data, next: nil}
	if q.IsEmpty() {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
	q.size++
	return q.size, nil
}

func (q *LinkedListQueue) DeQueue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, ErrQueueEmpty
	}
	removed := q.head.data
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return removed, nil
}

func (q *LinkedListQueue) Front() (interface{}, error) {
	if q.IsEmpty() {
		return nil, ErrQueueEmpty
	}
	return q.head.data, nil
}

func (q *LinkedListQueue) Rear() (interface{}, error) {
	if q.IsEmpty() {
		return nil, ErrQueueEmpty
	}
	return q.tail.data, nil
}

func (q *LinkedListQueue) Size() int {
	return q.size
}

func (q *LinkedListQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *LinkedListQueue) ToSlice() []interface{} {
	res := make([]interface{}, q.size)
	curr := q.head
	for i := 0; curr != nil; i++ {
		res[i] = curr.data
		curr = curr.next
	}
	return res
}
