package main

import "math"

func average(numbers []int) float64 {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return float64(sum) / float64(len(numbers))
}
func variance(numbers []int, average float64) float64 {
	var sum float64

	for _, number := range numbers {
		difference := float64(number) - average
		square := difference * difference

		sum += square
	}

	return sum / float64(len(numbers))
}
func standardDeviation(variance float64) float64 {
	return math.Sqrt(variance)
}
