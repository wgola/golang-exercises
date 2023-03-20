package challenge02

import (
	"math/big"
)

func fibbonaciCalculations(x int) (res map[int]int) {
	res = map[int]int{}

	helper := func(x int) {
		res[x] += 1
	}

	fibonacci(x, helper)

	return res
}

func fibonacci(x int, addOccurrance func(x int)) *big.Int {
	addOccurrance(x)
	if x == 2 || x == 1 {
		return new(big.Int).SetUint64(1)
	} else {
		return new(big.Int).Add(fibonacci(x-1, addOccurrance), fibonacci(x-2, addOccurrance))
	}
}
