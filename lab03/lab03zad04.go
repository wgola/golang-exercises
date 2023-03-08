package lab03

import (
	"fmt"
	"strconv"
)

func Zad04() {
	lengths := []rune{}

	for i := 1; i <= 100000; i++ {
		_, length := LongestCycle(i, i)
		letters := []rune(strconv.Itoa(length))
		lengths = append(lengths, letters...)
	}

	letters := make(map[string]int)

	for _, letter := range lengths {
		letters[string(letter)] += 1
	}

	mostInstances, mostNumbers := 0, ""

	for key, value := range letters {
		if mostInstances < value {
			mostInstances = value
			mostNumbers = key
		}

		fmt.Printf("Number: %v, instances: %v\n", key, value)
	}

	fmt.Printf("Number with most occurances: %v (%v)\n", mostNumbers, mostInstances)
}
