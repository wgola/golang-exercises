package lab04

import "fmt"

func zad04() {
	var (
		matrix1 = [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}}
		matrix2 = [][]int{{8, 7, 6}, {5, 4, 3}, {2, 1, 0}}
	)

	sum := addMatrix(matrix1, matrix2)
	fmt.Println("Matrix 1: ")
	printMatrix(matrix1)
	fmt.Println("Matrix 2: ")
	printMatrix(matrix2)
	fmt.Println("Sum of matrixes: ")
	printMatrix(sum)
}

func addMatrix(matrix1, matrix2 [][]int) (sumMatrix [][]int) {
	for i := 0; i < len(matrix1); i++ {
		row := []int{}
		for j := 0; j < len(matrix1[i]); j++ {
			row = append(row, matrix1[i][j]+matrix2[i][j])
		}
		sumMatrix = append(sumMatrix, row)
	}
	return
}

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%v ", matrix[i][j])
		}
		fmt.Println()
	}
}
