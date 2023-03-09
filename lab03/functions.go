package lab03

func longestCycle(start int, end int) (int, int) {
	var number, longest int

	collatz := func(number int) int {
		return helper(number, 0)
	}

	for i := start; i <= end; i++ {
		cycle := collatz(i)
		if longest < cycle {
			longest = cycle
			number = i
		}
	}

	return number, longest
}

func helper(number int, iterNum int) int {
	switch number {
	case 1:
		return iterNum
	default:
		if number%2 == 0 {
			return helper(number/2, iterNum+1)
		} else {
			return helper((number*3)+1, iterNum+1)
		}
	}
}
