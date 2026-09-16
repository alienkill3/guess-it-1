package main

import "math"

func predictRange(averageValue float64, stdDevValue float64) (int, int) {

	margin := stdDevValue * 0.1

	if margin < 1 {
		margin = 1
	}

	upperValue := math.Ceil(averageValue + margin)
	lowerValue := math.Floor(averageValue - margin)

	return int(lowerValue), int(upperValue)

}
