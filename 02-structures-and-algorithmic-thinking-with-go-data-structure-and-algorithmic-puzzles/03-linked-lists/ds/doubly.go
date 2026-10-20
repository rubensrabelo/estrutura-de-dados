package ds

type doublyNode struct {
	value int
	prev  *doublyNode
	next  *doublyNode
}

type DoublyLinkedList struct {
	head *doublyNode
	tail *doublyNode
	size int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (dll *DoublyLinkedList) Push(value int) int {
	newNode := &doublyNode{value: value}
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		dll.tail.next = newNode
		newNode.prev = dll.tail
		dll.tail = newNode
	}
	dll.size++
	return dll.size
}

func (dll *DoublyLinkedList) InsertAt(index int, value int) (int, error) {
	if index < 0 || index > dll.size {
		return 0, ErrIndexOutOfBounds
	}
	newNode := &doublyNode{value: value}
	if index == 0 {
		if dll.head == nil {
			dll.head = newNode
			dll.tail = newNode
		} else {
			newNode.next = dll.head
			dll.head.prev = newNode
			dll.head = newNode
		}
	} else if index == dll.size {
		dll.tail.next = newNode
		newNode.prev = dll.tail
		dll.tail = newNode
	} else {
		curr := dll.head
		for i := 0; i < index; i++ {
			curr = curr.next
		}
		newNode.prev = curr.prev
		newNode.next = curr
		curr.prev.next = newNode
		curr.prev = newNode
	}
	dll.size++
	return dll.size, nil
}

func (dll *DoublyLinkedList) DeleteAt(index int) (int, error) {
	if index < 0 || index >= dll.size {
		return 0, ErrIndexOutOfBounds
	}
	var removed int
	if index == 0 {
		removed = dll.head.value
		dll.head = dll.head.next
		if dll.head != nil {
			dll.head.prev = nil
		} else {
			dll.tail = nil
		}
	} else if index == dll.size-1 {
		removed = dll.tail.value
		dll.tail = dll.tail.prev
		if dll.tail != nil {
			dll.tail.next = nil
		} else {
			dll.head = nil
		}
	} else {
		curr := dll.head
		for i := 0; i < index; i++ {
			curr = curr.next
		}
		removed = curr.value
		curr.prev.next = curr.next
		curr.next.prev = curr.prev
	}
	dll.size--
	return removed, nil
}

func (dll *DoublyLinkedList) ToSlice() []int {
	res := make([]int, 0, dll.size)
	curr := dll.head
	for curr != nil {
		res = append(res, curr.value)
		curr = curr.next
	}
	return res
}
