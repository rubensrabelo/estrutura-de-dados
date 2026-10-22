package main

import (
	"dsproject/ds"
	"fmt"
)

func runNaryTreeRunner() {
	fmt.Println("--- N-ary Tree Operations (Generic Tree) ---")

	var naryTree ds.NaryTree = ds.NewSimpleNaryTree()

	// A primeira inserção define a Raiz. O parentValue (0) é ignorado.
	naryTree.Insert(0, 10) 

	// Inserindo filhos sob o nó raiz (10)
	naryTree.Insert(10, 20)
	naryTree.Insert(10, 30)
	naryTree.Insert(10, 40)

	// Inserindo netos sob o nó (20)
	naryTree.Insert(20, 50)
	naryTree.Insert(20, 60)

	// Inserindo neto sob o nó (30)
	naryTree.Insert(30, 70)

	naryTraverser := ds.NewNaryTraversalService(naryTree)

	fmt.Println("\n--- Executing N-ary Tree Traversals ---")
	fmt.Println("Pre-Order Traversal:  ", naryTraverser.PreOrder())
	fmt.Println("Post-Order Traversal: ", naryTraverser.PostOrder())
	fmt.Println("Level-Order (BFS):    ", naryTraverser.LevelOrder())

	fmt.Printf("\nN-ary Tree Size: %d\n", naryTree.Size())
	fmt.Printf("N-ary Tree Height: %d\n", naryTree.Height())

	fmt.Println("\n--- Testing N-ary Search ---")
	naryNode, err := naryTree.Search(60)
	if err == nil {
		fmt.Printf("Found N-ary node with value: %d\n", naryNode.Value)
	} else {
		fmt.Println("Error:", err)
	}

	fmt.Println("\n--- N-ary Auxiliary Operations ---")
	naryMaxLvl, naryMaxSum := naryTree.FindMaxSumLevel()
	fmt.Printf("Level with max sum: %d (Sum: %d)\n", naryMaxLvl, naryMaxSum)

	naryNode1, _ := naryTree.Search(50)
	naryNode2, _ := naryTree.Search(60)
	naryLcaNode := naryTree.LCA(naryNode1, naryNode2)
	if naryLcaNode != nil {
		fmt.Printf("LCA of 50 and 60 is: %d\n", naryLcaNode.Value)
	}

	naryNode3, _ := naryTree.Search(50)
	naryNode4, _ := naryTree.Search(70)
	naryLcaNode2 := naryTree.LCA(naryNode3, naryNode4)
	if naryLcaNode2 != nil {
		fmt.Printf("LCA of 50 and 70 is: %d\n", naryLcaNode2.Value)
	}

	fmt.Println("\n--- Testing N-ary Delete ---")
	naryDeleted, err := naryTree.Delete(20)
	if err == nil {
		fmt.Printf("Deleted value: %d\n", naryDeleted)
	} else {
		fmt.Println("Error:", err)
	}

	fmt.Println("Level-Order (BFS) After Delete:", naryTraverser.LevelOrder())
	fmt.Printf("N-ary Tree Size After Delete: %d\n", naryTree.Size())
}
