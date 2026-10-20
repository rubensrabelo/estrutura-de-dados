package ds

import "math/rand"

const maxLevel = 4

type skipNode struct {
	value int
	forward []*skipNode
}

type SkipList struct {
	head  *skipNode
	level int
}

func NewSkipList() *SkipList {
	return &SkipList{
		head:  &skipNode{value: -1, forward: make([]*skipNode, maxLevel)},
		level: 0,
	}
}

func (sl *SkipList) randomLevel() int {
	lvl := 0
	for rand.Float32() < 0.5 && lvl < maxLevel-1 {
		lvl++
	}
	return lvl
}

func (sl *SkipList) Push(value int) int {
	update := make([]*skipNode, maxLevel)
	curr := sl.head
	for i := sl.level; i >= 0; i-- {
		for curr.forward[i] != nil && curr.forward[i].value < value {
			curr = curr.forward[i]
		}
		update[i] = curr
	}
	lvl := sl.randomLevel()
	if lvl > sl.level {
		for i := sl.level + 1; i <= lvl; i++ {
			update[i] = sl.head
		}
		sl.level = lvl
	}
	newNode := &skipNode{value: value, forward: make([]*skipNode, lvl+1)}
	for i := 0; i <= lvl; i++ {
		newNode.forward[i] = update[i].forward[i]
		update[i].forward[i] = newNode
	}
	return 0
}

func (sl *SkipList) Find(value int) (int, error) {
	curr := sl.head
	for i := sl.level; i >= 0; i-- {
		for curr.forward[i] != nil && curr.forward[i].value < value {
			curr = curr.forward[i]
		}
	}
	curr = curr.forward[0]
	if curr != nil && curr.value == value {
		return curr.value, nil
	}
	return -1, ErrValueNotFound
}

func (sl *SkipList) ToSlice() []int {
	res := make([]int, 0)
	curr := sl.head.forward[0]
	for curr != nil {
		res = append(res, curr.value)
		curr = curr.forward[0]
	}
	return res
}
