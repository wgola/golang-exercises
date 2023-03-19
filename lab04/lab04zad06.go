package lab04

import "fmt"

func zad06() {
	matrix1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	matrix2 := [][]int{{9, 8, 7}, {6, 5, 4}, {3, 2, 1}}

	fmt.Println("Matrix 1: ")
	printMatrix(matrix1)
	fmt.Println("Matrix 2: ")
	printMatrix(matrix2)
	fmt.Println("Matrix 1 x Matrix 2: ")
	m1Xm2, _ := multiplyMatrix(matrix1, matrix2)
	printMatrix(m1Xm2)
	fmt.Println("Matrix 2 x Matrix 1: ")
	m2Xm1, _ := multiplyMatrix(matrix2, matrix1)
	printMatrix(m2Xm1)
}
