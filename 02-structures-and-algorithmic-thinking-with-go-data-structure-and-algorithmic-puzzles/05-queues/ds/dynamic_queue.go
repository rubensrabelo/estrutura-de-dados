package ds

type DynamicCircularArrayQueue struct {
	front    int
	rear     int
	size     int
	capacity int
	array    []interface{}
}

func NewDynamicCircularArrayQueue(capacity int) *DynamicCircularArrayQueue {
	if capacity <= 0 {
		capacity = 2
	}
	return &DynamicCircularArrayQueue{
		front:    0,
		rear:     -1,
		size:     0,
		capacity: capacity,
		array:    make([]interface{}, capacity),
	}
}

func (q *DynamicCircularArrayQueue) resize() {
	newCapacity := q.capacity * 2
	newArray := make([]interface{}, newCapacity)
	for i := 0; i < q.size; i++ {
		newArray[i] = q.array[(q.front+i)%q.capacity]
	}
	q.front = 0
	q.rear = q.size - 1
	q.capacity = newCapacity
	q.array = newArray
}

func (q *DynamicCircularArrayQueue) EnQueue(data interface{}) (int, error) {
	if q.size == q.capacity {
		q.resize()
	}
	q.rear = (q.rear + 1) % q.capacity
	q.array[q.rear] = data
	q.size++
	return q.size, nil
}

func (q *DynamicCircularArrayQueue) DeQueue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, ErrQueueEmpty
	}
	removed := q.array[q.front]
	q.array[q.front] = nil
	q.front = (q.front + 1) % q.capacity
	q.size--
	return removed, nil
}

func (q *DynamicCircularArrayQueue) Front() (interface{}, error) {
	if q.IsEmpty() {
		return nil, ErrQueueEmpty
	}
	return q.array[q.front], nil
}

func (q *DynamicCircularArrayQueue) Rear() (interface{}, error) {
	if q.IsEmpty() {
		return nil, ErrQueueEmpty
	}
	return q.array[q.rear], nil
}

func (q *DynamicCircularArrayQueue) Size() int {
	return q.size
}

func (q *DynamicCircularArrayQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *DynamicCircularArrayQueue) ToSlice() []interface{} {
	res := make([]interface{}, q.size)
	for i := 0; i < q.size; i++ {
		res[i] = q.array[(q.front+i)%q.capacity]
	}
	return res
}
