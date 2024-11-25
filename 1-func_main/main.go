package main

import "fmt"

func main() {

	// Slice declaration
	numbers := []int{1, 2, 3, 4, 5}

	// Create slice using make
	scores := make([]int, 5, 10) // length 5, capacity 10

	// Append to slice
	numbers = append(numbers, 6, 7, 8, 9, 10, 11)

	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Scores: %v\n", scores)

}
