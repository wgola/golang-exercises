package challenge02

import (
	"fmt"
)

func Challenge02() {
	var nick string

	fmt.Println("Enter your nick: ")
	fmt.Scanf("%v", &nick)

	strongNumber := findStrongNumber(nick)

	fact := factorial(strongNumber)

	fmt.Printf("Strong number: %v\n%v! = %v\n", strongNumber, strongNumber, fact)

	weakNumber, fibbonaciMap := findWeakNumber(strongNumber)

	fmt.Printf("Weak number: %v\nCalls of Fib(%v) = %v\n", weakNumber, weakNumber, fibbonaciMap[weakNumber])

}
