package main

import "fmt"

// Cenário recursivo
func factorial(n int) int {
	if n ==  0 {
		return 1
	}

	return n * factorial(n - 1)
}

// Cenário iterativo
func factorialIterative(n int) uint64 {
	var result uint64 = 1

	if n < 0 {
		fmt.Print("Não existe fatorial de numero negatio.")
	} else {
		for i:= 1; i<= n; i++ {
			result *= uint64(i)
		}
	}

	return result
}

func main() {
	fmt.Println("Recursivo: ", factorial(5))
	fmt.Println("Iterativo: ", factorialIterative(5))
}