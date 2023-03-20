package challenge02

import (
	"strconv"
	"strings"
)

func findStrongNumber(nick string) (strongNumber int64) {
	asciiTable := []string{}

	for _, letter := range nick {
		asciiTable = append(asciiTable, strconv.Itoa(int(letter)))
	}

	strongNumber = 1

MainLoop:
	for {
		fact := factorial(strongNumber).String()

		for _, ascii := range asciiTable {
			if !strings.Contains(fact, ascii) {
				strongNumber++
				continue MainLoop
			}
		}

		return
	}
}
