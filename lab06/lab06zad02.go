package lab06

import (
	"flag"
	"fmt"
	"math"
)

func z() {
	type T struct {
		a, b, c float64
	}
	a, b, c := *flag.Float64("a", 0, "a"), *flag.Float64("b", 0, "b"), *flag.Float64("c", 0, "c")
	flag.Parse()

	if r, s := func(t *T) (r T, s bool) {
		if d := t.b*t.b - 4*t.a*t.c; d >= 0 {
			s, r.a, r.b = true, (-t.b+math.Sqrt(d))/2*t.a, (-t.b-math.Sqrt(d))/2*t.a
		}
		return
	}(&T{a, b, c}); !s {
		fmt.Print("No answer\n")
	} else {
		fmt.Printf("R1=%v\nR2=%v\n", r.a, r.b)
	}
}
