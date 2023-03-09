package lab03

import "fmt"

func Zad04() {
	lengths := []int{}

	for i := 1; i <= 100_000; i++ {
		_, length := longestCycle(i, i)
		lengths = append(lengths, length)
	}

	stats := make(map[int]int)

	for i := range lengths {
		stats[lengths[i]] += 1
	}

	var biggestValue, foundKey int
	for key, value := range stats {
		if biggestValue < value {
			biggestValue = value
			foundKey = key
		}

		fmt.Printf("Key: %v, value: %v\n", key, value)
	}

	fmt.Printf("Biggest value: %v for key %v\n", biggestValue, foundKey)
}
