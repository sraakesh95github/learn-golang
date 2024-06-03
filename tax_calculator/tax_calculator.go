package main

import (
	"fmt"
)

func main() {

	earnings_before_tax, earnings_after_tax := 12000.0, 10000.0
	var ratio float64

	fmt.Print("Earnings before tax: ")
	fmt.Scan(&earnings_before_tax)

	fmt.Print("Earnings after tax: ")
	fmt.Scan(&earnings_after_tax)

	ratio = earnings_before_tax / earnings_after_tax

	fmt.Println("Ratio: ")
	fmt.Print(ratio)
}
