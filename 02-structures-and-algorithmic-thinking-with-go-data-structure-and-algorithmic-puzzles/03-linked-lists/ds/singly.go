package ds

type singlyNode struct {
	value int
	next  *singlyNode
}

type SinglyLinkedList struct {
	head *singlyNode
	size int
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

func (sll *SinglyLinkedList) Push(value int) int {
	newNode := &singlyNode{value: value}
	if sll.head == nil {
		sll.head = newNode
	} else {
		curr := sll.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = newNode
	}
	sll.size++
	return sll.size
}

func (sll *SinglyLinkedList) InsertAt(index int, value int) (int, error) {
	if index < 0 || index > sll.size {
		return 0, ErrIndexOutOfBounds
	}
	newNode := &singlyNode{value: value}
	if index == 0 {
		newNode.next = sll.head
		sll.head = newNode
	} else {
		curr := sll.head
		for i := 0; i < index-1; i++ {
			curr = curr.next
		}
		newNode.next = curr.next
		curr.next = newNode
	}
	sll.size++
	return sll.size, nil
}

func (sll *SinglyLinkedList) DeleteAt(index int) (int, error) {
	if index < 0 || index >= sll.size {
		return 0, ErrIndexOutOfBounds
	}
	var removed int
	if index == 0 {
		removed = sll.head.value
		sll.head = sll.head.next
	} else {
		curr := sll.head
		for i := 0; i < index-1; i++ {
			curr = curr.next
		}
		removed = curr.next.value
		curr.next = curr.next.next
	}
	sll.size--
	return removed, nil
}

func (sll *SinglyLinkedList) Find(value int) (int, error) {
	curr := sll.head
	for i := 0; i < sll.size; i++ {
		if curr.value == value {
			return i, nil
		}
		curr = curr.next
	}
	return -1, ErrValueNotFound
}

func (sll *SinglyLinkedList) ToSlice() []int {
	res := make([]int, 0, sll.size)
	curr := sll.head
	for curr != nil {
		res = append(res, curr.value)
		curr = curr.next
	}
	return res
}
