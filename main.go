package main

import "os"

func main() {

	err := readInput()
	if err != nil {
		os.Exit(1)
	}
}
