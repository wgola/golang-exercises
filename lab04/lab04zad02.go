package lab04

import "fmt"

func zad02() {
	tab1, tab2 := [5]int{}, [5]int{}

	for i := 0; i < len(tab1); i++ {
		tab1[i], tab2[i] = i, i
	}
	result := hadamardProduct(tab1[:], tab2[:])

	fmt.Printf("Array 1: %v\nArray 2: %v\nHadamard's product: %v\n", tab1, tab2, result)
}

func hadamardProduct(a, b []int) (product []int) {
	for i := 0; i < len(a); i++ {
		product = append(product, a[i]*b[i])
	}
	return
}
