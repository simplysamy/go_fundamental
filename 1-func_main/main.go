package main

import (
	"fmt"
	"reflect"
)

// Struct can be compared as class concept in OOPS

// Define a struct
type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	// Create struct instance
	person := Person{
		Name:    "John Doe",
		Age:     30,
		Address: "123 Main St",
	}

	fmt.Printf("Person: %+v\n", person)

	t := reflect.TypeOf(person)
	fmt.Println(t) // Output
}
