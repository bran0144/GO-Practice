package main

import (
	"fmt"
	"math"
)

func maximum(numbers ...float64) float64 {
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}
func main() {
	fmt.Println(maximum(71, 56.2, 89.5, 102.3))
}
