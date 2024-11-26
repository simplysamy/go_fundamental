package main

import "fmt"

func main() {
	x := 42
	var ptr *int = &x // Pointer to int

	fmt.Printf("Value of x: %v\n", x)
	fmt.Printf("Address of x: %v\n", ptr)
	fmt.Printf("Value at pointer: %v\n", *ptr)
}
