package lab04

import "fmt"

func zad04() {
	matrix1, matrix2 := generateMatrix(10), generateMatrix(10)

	sum := addMatrix(matrix1, matrix2)
	fmt.Println("Matrix 1: ")
	printMatrix(matrix1)
	fmt.Println("Matrix 2: ")
	printMatrix(matrix2)
	fmt.Println("Sum of matrixes: ")
	printMatrix(sum)
}

func generateMatrix(size int) (matrix [][]int) {
	for i := 0; i < size; i++ {
		row := []int{}
		for j := 0; j < size; j++ {
			row = append(row, j+i*size)
		}
		matrix = append(matrix, row)
	}
	return
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
			fmt.Printf("%3d ", matrix[i][j])
		}
		fmt.Println()
	}
}
