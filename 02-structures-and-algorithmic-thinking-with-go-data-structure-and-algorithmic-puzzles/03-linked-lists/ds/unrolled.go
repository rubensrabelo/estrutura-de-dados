package ds

const nodeCapacity = 3

type unrolledNode struct {
	values []int
	numElements int
	next   *unrolledNode
}

type UnrolledLinkedList struct {
	head *unrolledNode
	size int
}

func NewUnrolledLinkedList() *UnrolledLinkedList {
	return &UnrolledLinkedList{}
}

func (ull *UnrolledLinkedList) Push(value int) int {
	if ull.head == nil {
		ull.head = &unrolledNode{values: make([]int, nodeCapacity)}
		ull.head.values[0] = value
		ull.head.numElements = 1
	} else {
		curr := ull.head
		for curr.next != nil {
			curr = curr.next
		}
		if curr.numElements < nodeCapacity {
			curr.values[curr.numElements] = value
			curr.numElements++
		} else {
			newNode := &unrolledNode{values: make([]int, nodeCapacity)}
			newNode.values[0] = value
			newNode.numElements = 1
			curr.next = newNode
		}
	}
	ull.size++
	return ull.size
}

func (ull *UnrolledLinkedList) ToSlice() []int {
	res := make([]int, 0, ull.size)
	curr := ull.head
	for curr != nil {
		for i := 0; i < curr.numElements; i++ {
			res = append(res, curr.values[i])
		}
		curr = curr.next
	}
	return res
}
