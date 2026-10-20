package main

import (
	"dsproject/ds"
	"fmt"
)

func main() {
	fmt.Println("=== 1. Simple Circular Array Queue ===")
	cq := ds.NewCircularArrayQueue(2)
	_, _ = cq.EnQueue(10)
	_, _ = cq.EnQueue(20)
	_, err := cq.EnQueue(30)
	if err != nil {
		fmt.Printf("Expected Overflow: %v\n", err)
	}
	fmt.Println("Queue Data:", cq.ToSlice())
	val, _ := cq.DeQueue()
	fmt.Printf("DeQueued: %v, New Front: ", val)
	_, _ = cq.EnQueue(30)
	fmt.Println(cq.ToSlice())

	fmt.Println("\n=== 2. Dynamic Circular Array Queue ===")
	dq := ds.NewDynamicCircularArrayQueue(2)
	_, _ = dq.EnQueue("A")
	_, _ = dq.EnQueue("B")
	_, _ = dq.EnQueue("C") // Força redimensionamento automático de 2 para 4
	fmt.Println("Dynamic Queue Data:", dq.ToSlice())
	_, _ = dq.DeQueue()
	fmt.Println("After DeQueue:", dq.ToSlice())

	fmt.Println("\n=== 3. Linked List Queue ===")
	lq := ds.NewLinkedListQueue()
	_, _ = lq.EnQueue(100)
	_, _ = lq.EnQueue(200)
	fmt.Println("Linked Queue Data:", lq.ToSlice())
	_, _ = lq.DeQueue()
	fmt.Println("After DeQueue:", lq.ToSlice())
}
