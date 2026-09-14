package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readInput() {
	scanner := bufio.NewScanner(os.Stdin)
	numbers := []int{}

	for scanner.Scan() {
		line := scanner.Text()

		number, err := strconv.Atoi(line)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid input. Please enter a valid integer.")
			return
		}
		numbers = append(numbers, number)
		if len(numbers) > 3 {
			numbers = numbers[1:]
		}
		fmt.Println(numbers)

	}
}
