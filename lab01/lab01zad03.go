package lab01

import (
	"fmt"
	"math"
)

func Zad03() {
	fmt.Println("Enter parameters of quadratic polynomial separated with space (a b c): ")

	var a, b, c float64
	fmt.Scanf("%f %f %f", &a, &b, &c)

	delta := delta(a, b, c)

	if delta == 0 {
		x1 := (-b - math.Sqrt(delta)) / 2 * a

		fmt.Printf("The solution is: %3.3f\n", x1)
	} else if delta > 0 {
		x1, x2 := (-b-math.Sqrt(delta))/2*a, (-b+math.Sqrt(delta))/2*a

		fmt.Printf("The solutions are: %3.3f and %3.3f\n", x1, x2)
	} else {
		imag := delta / -1
		sqrt := complex(0, math.Sqrt(imag))
		x1, x2 := (-complex(b, 0)-sqrt)/complex(2*a, 0), (-complex(b, 0)+sqrt)/complex(2*a, 0)

		fmt.Printf("The solutions are: %3.3f and %3.3f\n", x1, x2)
	}
}

func delta(a float64, b float64, c float64) float64 {
	return math.Pow(b, 2) - 4*a*c
}
