package ds

type circularNode struct {
	value int
	next  *circularNode
}

type CircularLinkedList struct {
	tail *circularNode
	size int
}

func NewCircularLinkedList() *CircularLinkedList {
	return &CircularLinkedList{}
}

func (cll *CircularLinkedList) Push(value int) int {
	newNode := &circularNode{value: value}
	if cll.tail == nil {
		newNode.next = newNode
		cll.tail = newNode
	} else {
		newNode.next = cll.tail.next
		cll.tail.next = newNode
		cll.tail = newNode
	}
	cll.size++
	return cll.size
}

func (cll *CircularLinkedList) DeleteAt(index int) (int, error) {
	if index < 0 || index >= cll.size {
		return 0, ErrIndexOutOfBounds
	}
	var removed int
	if cll.size == 1 {
		removed = cll.tail.value
		cll.tail = nil
	} else if index == 0 {
		head := cll.tail.next
		removed = head.value
		cll.tail.next = head.next
	} else {
		curr := cll.tail.next
		for i := 0; i < index-1; i++ {
			curr = curr.next
		}
		removed = curr.next.value
		if curr.next == cll.tail {
			cll.tail = curr
		}
		curr.next = curr.next.next
	}
	cll.size--
	return removed, nil
}

func (cll *CircularLinkedList) ToSlice() []int {
	if cll.tail == nil {
		return []int{}
	}
	res := make([]int, 0, cll.size)
	head := cll.tail.next
	curr := head
	for {
		res = append(res, curr.value)
		curr = curr.next
		if curr == head {
			break
		}
	}
	return res
}
