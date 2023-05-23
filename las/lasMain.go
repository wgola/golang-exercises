package las

import (
	"fmt"
	"math"
)

func LasMain() {
	allResults := make([]float64, 0)
	lastMean := 0.0

	for {
		forest := CreateNewForest(10, 10)

		forest.RandomizeForest(0.7)

		allTrees := forest.GetNumberOfTrees()

		forest.RandomizeThunderStrike()

		if numberOfFiredTrees := forest.GetNumberOfBurnedTrees(); numberOfFiredTrees != 0 {
			for {
				forest.SimulateTurn()
				newNumber := forest.GetNumberOfBurnedTrees()
				if newNumber == numberOfFiredTrees {
					break
				}
				numberOfFiredTrees = newNumber
			}
		}

		burnedTrees := forest.GetNumberOfBurnedTrees()

		allResults = append(allResults, float64(burnedTrees)/float64(allTrees))
		currentMean := mean(allResults)

		if math.Abs(currentMean-lastMean) < 0.01 {
			lastMean = currentMean
			break
		}

		lastMean = currentMean
	}

	fmt.Println(lastMean)
}

func mean(numbers []float64) (mean float64) {
	for _, n := range numbers {
		mean += n
	}

	mean = mean / float64(len(numbers))
	return
}
