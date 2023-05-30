package lab08

import (
	"fmt"
	"math"
)

type resultsList []float64

func LasMain() {
	mapOfResuls := make(map[string]float64, 0)
	for i := 5; i <= 250; i += 5 {
		result := simulateForBoardSize(i, i)
		mapOfResuls[fmt.Sprintf("%dx%d", i, i)] = result

		fmt.Printf("%dx%d = %.2f%%\n", i, i, result*100)
	}
}

func simulateForBoardSize(x, y int) float64 {
	var allResults resultsList = make([]float64, 0)
	lastMean := 0.0

	for {
		forest := CreateNewForest(x, y)

		forest.RandomizeForest(0.5)

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

		if burnedTrees := forest.GetNumberOfBurnedTrees(); burnedTrees != 0 {
			allResults = append(allResults, float64(burnedTrees)/float64(allTrees))
			currentMean := allResults.mean()

			if math.Abs(currentMean-lastMean) < 0.01 {
				lastMean = currentMean
				break
			}

			lastMean = currentMean
		}
	}

	return lastMean
}

func (r *resultsList) mean() (mean float64) {
	for _, n := range *r {
		mean += n
	}

	mean = mean / float64(len(*r))
	return
}
