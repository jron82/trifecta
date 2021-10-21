package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

type roots struct {
	first  float64
	second float64
}

func calc(a int, b int, c int) roots {
	negativeB := float64(-b)

	precursor := float64((b * b) - (4 * a * c))
	discrimant := math.Sqrt(math.Abs((precursor)))

	denominator := float64(2 * a)

	first := (negativeB + discrimant) / denominator
	second := (negativeB - discrimant) / denominator

	return roots{first: first, second: second}

}

func main() {
	a := flag.Int("a", 0, "The coefficient in Ax^2")
	b := flag.Int("b", 0, "The coefficient in bx")
	c := flag.Int("c", 0, "The constant c")

	flag.Parse()

	if *a == 0 || *b == 0 || *c == 0 {
		fmt.Println("Missing command. Can't evaluate a trinomial without 3 args :-)")
		os.Exit(2)
	}

	res := calc(*a, *b, *c)

	text := fmt.Sprintf("The roots are %.2f and %.2f", res.first, res.second)
	fmt.Println(text)

}
