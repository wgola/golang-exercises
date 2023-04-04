package lab06

import (
	"flag"
	"fmt"
	"math"
)

type Trinomial struct {
	a, b, c float64
}

type Tuple struct {
	a, b float64
}

func zad01() {
	var a, b, c float64
	flag.Float64Var(&a, "a", 0, "Parameter 'a'")
	flag.Float64Var(&b, "b", 0, "Parameter 'b'")
	flag.Float64Var(&c, "c", 0, "Parameter 'c'")
	flag.Parse()

	trinomial := &Trinomial{a, b, c}

	if result, status := countRoots(trinomial); !status {
		fmt.Println("There are no square roots of this trinomial!")
	} else {
		fmt.Printf("Root 1: %v\nRoot 2: %v\n", result.a, result.b)
	}
}

func countRoots(t *Trinomial) (res Tuple, status bool) {
	if delta := math.Pow(t.b, 2) - 4*t.a*t.c; delta < 0 {
		status = false
	} else {
		status = true
		res.a = (-t.b + math.Sqrt(delta)) / 2 * t.a
		res.b = (-t.b - math.Sqrt(delta)) / 2 * t.a
	}
	return
}
