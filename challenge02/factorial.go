package challenge02

import "math/big"

func factorial(x int64) (fact *big.Int) {
	fact = new(big.Int).MulRange(1, x)
	return
}
