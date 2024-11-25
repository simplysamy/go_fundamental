package main

import "fmt"

func main() {
	// Map declaration
	studentScores := map[string]int{
		"John":  95,
		"Alice": 89,
		"Bob":   78,
	}

	// Create map using make
	ages := make(map[string]int)

	// Add key-value pairs
	ages["Tom"] = 25
	ages["Jane"] = 28

	fmt.Printf("Student Scores: %v\n", studentScores)
	fmt.Printf("Ages: %v\n", ages)
}
