// NOTA: O Garbage Collector do Go proíbe aritmética de ponteiros insegura, por isso uma simulação de arena/tabela é usada para preservar fielmente a lógica XOR de eficiência de memória.
package ds

type xorNode struct {
	value int
	diff  int
}

type XORLinkedList struct {
	arena []xorNode
	head  int
	tail  int
	size  int
}

func NewXORLinkedList() *XORLinkedList {
	return &XORLinkedList{
		arena: make([]xorNode, 0),
		head:  -1,
		tail:  -1,
	}
}

func (xll *XORLinkedList) Push(value int) int {
	newNode := xorNode{value: value, diff: 0}
	xll.arena = append(xll.arena, newNode)
	newIdx := len(xll.arena) - 1

	if xll.head == -1 {
		xll.head = newIdx
		xll.tail = newIdx
	} else {
		xll.arena[xll.tail].diff = xll.arena[xll.tail].diff ^ newIdx
		xll.arena[newIdx].diff = xll.tail ^ (-1)
		xll.tail = newIdx
	}
	xll.size++
	return xll.size
}

func (xll *XORLinkedList) ToSlice() []int {
	res := make([]int, 0, xll.size)
	curr := xll.head
	prev := -1
	for curr != -1 && len(res) < xll.size {
		res = append(res, xll.arena[curr].value)
		next := xll.arena[curr].diff ^ prev
		prev = curr
		curr = next
	}
	return res
}
