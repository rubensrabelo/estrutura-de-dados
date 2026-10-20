// Nota: Eu ainda vou reformular as estruturas de dados (Releitura do Livro)
package main

import (
	"dsproject/ds"
	"fmt"
)

func main() {
	fmt.Println("--- 3.4/3.5 ArrayList ---")
	al := ds.NewArrayList(5)
	_, _ = al.Push(10)
	_, _ = al.Push(20)
	_, _ = al.InsertAt(1, 15)
	fmt.Println("ArrayList Data:", al.ToSlice())

	fmt.Println("\n--- 3.6 Singly Linked List ---")
	sll := ds.NewSinglyLinkedList()
	sll.Push(10)
	sll.Push(20)
	_, _ = sll.InsertAt(1, 15)
	fmt.Println("Singly Data:", sll.ToSlice())

	fmt.Println("\n--- 3.7 Doubly Linked List ---")
	dll := ds.NewDoublyLinkedList()
	dll.Push(10)
	dll.Push(20)
	_, _ = dll.InsertAt(1, 15)
	fmt.Println("Doubly Data:", dll.ToSlice())

	fmt.Println("\n--- 3.8 Circular Linked List ---")
	cll := ds.NewCircularLinkedList()
	cll.Push(10)
	cll.Push(20)
	cll.Push(30)
	fmt.Println("Circular Data:", cll.ToSlice())

	fmt.Println("\n--- 3.9 XOR Linked List ---")
	xll := ds.NewXORLinkedList()
	xll.Push(10)
	xll.Push(20)
	xll.Push(30)
	fmt.Println("XOR Data:", xll.ToSlice())

	fmt.Println("\n--- 3.10 Unrolled Linked List ---")
	ull := ds.NewUnrolledLinkedList()
	ull.Push(1)
	ull.Push(2)
	ull.Push(3)
	ull.Push(4)
	fmt.Println("Unrolled Data:", ull.ToSlice())

	fmt.Println("\n--- 3.11 Skip List ---")
	sl := ds.NewSkipList()
	sl.Push(30)
	sl.Push(10)
	sl.Push(20)
	fmt.Println("SkipList (Sorted Output):", sl.ToSlice())
	
	val, err := sl.Find(20)
	if err == nil {
		fmt.Printf("Search successful! Found value: %d\n", val)
	}
}
