package main

func main() {
	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10

	//change investmentAmount to float64
	//expectedReturnRate is in float64

	var futureValue = float64(investmentAmount) * (1 + expectedReturnRate/100) * *float64(years)
}
