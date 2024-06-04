package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {

	const inflationRate float64 = 6.5
	years, expectedReturnRate := 10.0, 5.5
	var investmentAmount float64

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	//Simple string formatting
	// fmt.Printf("Future value: %.1f\nFuture real value: %.1f\n\n", futureValue, futureRealValue)

	//String returns
	formattedFV := fmt.Sprintf("Future value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future real value: %.1f\n\n", futureRealValue)

	print(formattedFV + formattedRFV)
}

func outputText(text string) {
	fmt.Println(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
