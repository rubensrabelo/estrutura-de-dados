package main

import (
	"dsproject/ds"
	"fmt"
)

func main() {
	// Inicializa a estrutura usando a interface
	var uf ds.DisjointSet = ds.NewUnionFind()

	fmt.Println("--- 1. Inicializando Elementos (MakeSet) ---")
	elements := []int{10, 20, 30, 40, 50}
	for _, el := range elements {
		uf.MakeSet(el)
	}
	fmt.Printf("Total de elementos (Size): %d\n", uf.Size())
	fmt.Printf("Total de conjuntos iniciais (CountSets): %d\n\n", uf.CountSets())

	fmt.Println("--- 2. Testando Uniões (Union) ---")
	// Une 10 e 20 -> Conjunto {10, 20}
	merged, _ := uf.Union(10, 20)
	fmt.Printf("Unindo 10 e 20: %t\n", merged)

	// Une 30 e 40 -> Conjunto {30, 40}
	uf.Union(30, 40)
	
	// Une 20 e 30 -> Conecta os dois grupos -> Conjunto {10, 20, 30, 40}
	merged, _ = uf.Union(20, 30)
	fmt.Printf("Unindo 20 e 30 (mesclando grupos): %t\n", merged)

	// Tenta unir elementos que já pertencem ao mesmo grupo
	alreadyMerged, _ := uf.Union(10, 40)
	fmt.Printf("Tentando unir 10 e 40 novamente: %t\n\n", alreadyMerged)

	fmt.Println("--- 3. Verificando Representantes (Find) ---")
	rep10, _ := uf.Find(10)
	rep40, _ := uf.Find(40)
	rep50, _ := uf.Find(50) // Ficou isolado
	fmt.Printf("Representante do elemento 10: %d\n", rep10)
	fmt.Printf("Representante do elemento 40: %d\n", rep40)
	fmt.Printf("Representante do elemento 50: %d\n", rep50)
	fmt.Printf("10 e 40 pertencem ao mesmo conjunto? %t\n\n", rep10 == rep40)

	fmt.Println("--- 4. Estado Final ---")
	fmt.Printf("Total de elementos: %d\n", uf.Size())
	fmt.Printf("Total de conjuntos disjuntos restantes: %d\n\n", uf.CountSets())

	fmt.Println("--- 5. Testando Tratamento de Erros ---")
	_, err := uf.Find(999)
	if err != nil {
		fmt.Printf("Erro esperado ao buscar 999: %v\n", err)
	}
}
