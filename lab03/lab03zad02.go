package lab03

import "fmt"

func Zad02() {
	start := 1000
	for i := 1; i < 100; i++ {
		number, length := LongestCycle(start, start+1000)
		fmt.Printf("Number: %v, number of cycles: %v.\n", number, length)
		start += i * 1000
	}
}
