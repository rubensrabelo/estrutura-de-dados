package main

import (
	"dsproject/ds"
	"fmt"
)

func main() {
	fmt.Println("--- 6.5 Binary Tree Operations (Recursive) ---")
	
	var tree ds.BinaryTree = ds.NewSimpleBinaryTree()

	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)

	traverser := ds.NewTraversalService(tree)

	fmt.Println("\n--- Executing Tree Traversals (via DI) ---")
	fmt.Println("Pre-Order Traversal: ", traverser.PreOrder())
	fmt.Println("In-Order Traversal:  ", traverser.InOrder())
	fmt.Println("Post-Order Traversal:", traverser.PostOrder())
	fmt.Println("Level-Order (BFS):   ", traverser.LevelOrder())

	fmt.Printf("\nTree Size: %d\n", tree.Size())
	fmt.Printf("Tree Height: %d\n", tree.Height())

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
	lcaNode := tree.LCA(node1, node2)
	if lcaNode != nil {
		fmt.Printf("LCA of 4 and 5 is: %d\n", lcaNode.Value)
	}

	fmt.Println("\n--- Testing Delete ---")
	deleted, err := tree.Delete(2)
	if err == nil {
		fmt.Printf("Deleted value: %d\n", deleted)
	}

	fmt.Println("In-Order Traversal After Delete:", traverser.InOrder())
}
