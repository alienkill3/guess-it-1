package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readInput() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid input. Please enter a valid integer.")
			return
		}
		fmt.Println(number)

	}
}
