package lab03

import "fmt"

func Zad01() {
	fmt.Println("Insert start and end of comparment separated with space (only natural numbers): ")
	var (
		start int
		end   int
	)
	fmt.Scanf("%v %v", &start, &end)
	number, length := longestCycle(start, end)
	fmt.Printf("Number: %v, number of cycles: %v.\n", number, length)
}
