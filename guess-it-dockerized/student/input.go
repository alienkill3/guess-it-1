package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readInput() error {
	scanner := bufio.NewScanner(os.Stdin)
	numbers := []int{}

	for scanner.Scan() {
		line := scanner.Text()

		number, err := strconv.Atoi(line)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid input. Please enter a valid integer.")
			return err
		}
		numbers = append(numbers, number)
		if len(numbers) > 3 {
			numbers = numbers[1:]
		}
		averageValue := average(numbers)
		varianceValue := variance(numbers, averageValue)
		stdDevValue := standardDeviation(varianceValue)
		lower, upper := predictRange(averageValue, stdDevValue)
		fmt.Println(lower, upper)
	}
	err := scanner.Err()

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		return err
	}
	return nil
}
