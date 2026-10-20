package ds

import "errors"

var (
	ErrIndexOutOfBounds = errors.New("ds: index out of bounds")
	ErrValueNotFound    = errors.New("ds: value not found")
	ErrEmptyList        = errors.New("ds: list is empty")
	ErrCapacityFull     = errors.New("ds: array list capacity full")
)

type ArrayList struct {
	data     []int
	size     int
	capacity int
}

func NewArrayList(capacity int) *ArrayList {
	return &ArrayList{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (al *ArrayList) Push(value int) (int, error) {
	if al.size >= al.capacity {
		return 0, ErrCapacityFull
	}
	al.data[al.size] = value
	al.size++
	return al.size, nil
}

func (al *ArrayList) InsertAt(index int, value int) (int, error) {
	if index < 0 || index > al.size {
		return 0, ErrIndexOutOfBounds
	}
	if al.size >= al.capacity {
		return 0, ErrCapacityFull
	}
	copy(al.data[index+1:al.size+1], al.data[index:al.size])
	al.data[index] = value
	al.size++
	return al.size, nil
}

func (al *ArrayList) DeleteAt(index int) (int, error) {
	if index < 0 || index >= al.size {
		return 0, ErrIndexOutOfBounds
	}
	removed := al.data[index]
	copy(al.data[index:al.size-1], al.data[index+1:al.size])
	al.size--
	al.data[al.size] = 0
	return removed, nil
}

func (al *ArrayList) UpdateAt(index int, value int) (int, error) {
	if index < 0 || index >= al.size {
		return 0, ErrIndexOutOfBounds
	}
	oldValue := al.data[index]
	al.data[index] = value
	return oldValue, nil
}

func (al *ArrayList) Find(value int) (int, error) {
	for i := 0; i < al.size; i++ {
		if al.data[i] == value {
			return i, nil
		}
	}
	return -1, ErrValueNotFound
}

func (al *ArrayList) Get(index int) (int, error) {
	if index < 0 || index >= al.size {
		return 0, ErrIndexOutOfBounds
	}
	return al.data[index], nil
}

func (al *ArrayList) Size() int {
	return al.size
}

func (al *ArrayList) Capacity() int {
	return al.capacity
}

func (al *ArrayList) Clear() {
	for i := 0; i < al.size; i++ {
		al.data[i] = 0
	}
	al.size = 0
}

func (al *ArrayList) ToSlice() []int {
	return al.data[:al.size]
}
