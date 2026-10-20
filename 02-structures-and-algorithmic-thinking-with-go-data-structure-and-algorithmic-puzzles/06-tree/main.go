package main

import (
	"dsproject/ds"
	"fmt"
)

func main() {
	fmt.Println("--- 6.5 Binary Tree Operations (Recursive) ---")
	tree := ds.NewBinaryTree()

	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)

	var elements []int
	tree.TraverseInOrder(tree.Root, &elements)
	fmt.Println("In-Order Traversal:", elements)
	fmt.Printf("Tree Size: %d\n", tree.Size(tree.Root))
	fmt.Printf("Tree Height: %d\n", tree.Height(tree.Root))

	fmt.Println("\n--- Testing Search ---")
	node, err := tree.Search(3)
	if err == nil {
		fmt.Printf("Found node with value: %d\n", node.Value)
	}

	fmt.Println("\n--- Auxiliary Operations ---")
	maxLvl, maxSum := tree.FindMaxSumLevel()
	fmt.Printf("Level with max sum: %d (Sum: %d)\n", maxLvl, maxSum)

	node1, _ := tree.Search(4)
	node2, _ := tree.Search(5)
	lcaNode := tree.LCA(tree.Root, node1, node2)
	if lcaNode != nil {
		fmt.Printf("LCA of 4 and 5 is: %d\n", lcaNode.Value)
	}

	fmt.Println("\n--- Testing Delete ---")
	deleted, err := tree.Delete(2)
	if err == nil {
		fmt.Printf("Deleted value: %d\n", deleted)
	}

	elements = []int{}
	tree.TraverseInOrder(tree.Root, &elements)
	fmt.Println("In-Order Traversal After Delete:", elements)
}
