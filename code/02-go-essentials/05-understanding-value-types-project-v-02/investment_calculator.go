package main

import "math"

func main() {
	var investmentAmount = 1000  // int → whole number
	var expectedReturnRate = 5.5 // float64 → decimal number
	var years = 10               // int

	var futureValue = float64(investmentAmount) * math.Pow(1+
		expectedReturnRate/100, float64(years))
}
