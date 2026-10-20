package main

import (
	"dsproject/ds"
	"fmt"
)

func main() {
	fmt.Println("=== 4.7 Simple Array Stack Implementation ===")
	arrayStack := ds.NewStack(3)

	_, _ = arrayStack.Push(10)
	_, _ = arrayStack.Push(20)
	_, _ = arrayStack.Push(30)
	fmt.Print("Array Stack Data: ")
	fmt.Println(arrayStack.ToSlice())
	fmt.Printf("Array Stack Size: %d\n", arrayStack.Size())

	valArr, _ := arrayStack.Pop()
	fmt.Printf("Popped from Array Stack: %v\n", valArr)
	fmt.Print("Array Stack Data After Pop: ")
	fmt.Println(arrayStack.ToSlice())
	fmt.Printf("Array Stack Size After Pop: %d\n", arrayStack.Size())

	fmt.Println("\n=== 4.9 Linked List Stack Implementation ===")
	llStack := ds.NewLinkedListStack()

	_, _ = llStack.Push("A")
	_, _ = llStack.Push("B")
	_, _ = llStack.Push("C")
	fmt.Print("Linked List Stack Data (Top to Bottom): ")
	fmt.Println(llStack.ToSlice())
	fmt.Printf("Linked List Stack Size: %d\n", llStack.Size())

	topVal, _ := llStack.Peek()
	fmt.Printf("Top element (Peek): %v\n", topVal)

	valLL, _ := llStack.Pop()
	fmt.Printf("Popped from Linked List Stack: %v\n", valLL)
	fmt.Print("Linked List Stack Data After Pop: ")
	fmt.Println(llStack.ToSlice())
	fmt.Printf("Linked List Stack Size After Pop: %d\n", llStack.Size())
}
