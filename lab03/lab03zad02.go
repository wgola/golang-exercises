package lab03

import "fmt"

func Zad02() {
	start := 1000
	for i := 1; i < 100; i++ {
		number, length := LongestCycle(start, start+1000)
		fmt.Printf("Range: %v-%v, number: %v, number of cycles: %v.\n", start, start+1000, number, length)
		start += 1000
	}
}
