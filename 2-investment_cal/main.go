package main

import (
	"fmt"
	"math"
)

func main() {
	// var investmentAmount float64 = 1000
	// var years float64 = 10
	var investmentAmount, years float64 = 1000, 10 // Multiple declaration

	expectedReturnRate := 5.5

	futureValue := (investmentAmount) * math.Pow(1+expectedReturnRate/100, (years))
	fmt.Println(futureValue)
}
