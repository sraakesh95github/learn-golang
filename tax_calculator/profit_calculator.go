package main

import (
	"fmt"
)

func main() {

	var tax_rate float64
	var revenue float64
	var expenses float64

	revenue = getUserInput("Revenue: ")
	expenses = getUserInput("Expenses: ")
	tax_rate = getUserInput("Tax Rate: ")

	earnings_before_tax, profit, ratio := calculateFinancials(revenue, expenses, tax_rate)

	fmt.Printf("Earnings before tax: %.2f\n", earnings_before_tax)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
}

func calculateFinancials(revenue, expenses, tax_rate float64) (float64, float64, float64) {
	earnings_before_tax := revenue - expenses
	profit := revenue - (revenue * (tax_rate / 100))
	ratio := earnings_before_tax / profit

	return earnings_before_tax, profit, ratio
}

func getUserInput(infoText string) float64 {

	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)

	return userInput
}
