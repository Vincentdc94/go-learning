package main

import (
	"fmt"
	"math"
)

//Dit is de Sqrt functie om te maken
func Sqrt(x float64) float64 {
	z := float64(2.)
	s := float64(0)

	for {
		z = z - (z*z-x)/(2*z)
		if math.Abs(s-z) < 1e-15 {
			break
		}
		s = z
	}

	return s
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
