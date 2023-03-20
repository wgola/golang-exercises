package challenge02

import "math"

func findWeakNumber(strongNumber int64) (weakNumber int, fibbonaciMap map[int]int) {
	fibbonaciMap = fibbonaciCalculations(30)

	closestValue := abs(fibbonaciMap[1] - int(strongNumber))
	for key, value := range fibbonaciMap {
		testValue := abs(value - int(strongNumber))
		if testValue < closestValue {
			closestValue = testValue
			weakNumber = key
		}
	}

	return
}

func abs(x int) (res int) {
	res = int(math.Abs(float64(x)))
	return
}
