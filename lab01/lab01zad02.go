package main

import "fmt"

func main() {
	fmt.Println("How old are you?")

	var i int
	_, err := fmt.Scanf("%d", &i)

	if err != nil {
		fmt.Println("Error - wrong age!")
		return
	}

	fmt.Println("Your age [in years] on different planets:")
	fmt.Printf("Mercury: %10f\n", float32(i)*0.24)
	fmt.Printf("Venus: %10f\n", float32(i)*0.62)
	fmt.Printf("Mars: %10f\n", float32(i)*1.88)
	fmt.Printf("Jupiter: %10f\n", float32(i)*11.87)
	fmt.Printf("Saturn: %10f\n", float32(i)*29.48)
	fmt.Printf("Uranus: %10f\n", float32(i)*84.07)
	fmt.Printf("Neptun: %10f\n", float32(i)*164.9)
}
