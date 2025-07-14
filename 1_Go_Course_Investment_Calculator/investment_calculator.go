package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64 = 1000
	rateOfReturn := 5.5
	var years float64
	// super short notation alternative
	//investmentAmount, years, rateOfReturn := 1000.0, 10.0, 5.5 // initial investment, years to grow, rate of return
	//fmt.Print("Enter the investment amount: ")
	outputText("Enter the investment amount:")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the rate of return: ")
	fmt.Scan(&rateOfReturn)

	fmt.Print("Enter the years value: ")
	fmt.Scan(&years)

	//futureValue := investmentAmount * math.Pow(1+rateOfReturn/100, years) // types mater when calculating the final value
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years) // future value adjusted for inflation
	futureValue, futureRealValue := calculateFutureValues(investmentAmount, rateOfReturn, years)

	// FORMAT OPTION 1:
	println("\n\nOPTION 1:")
	//fmt.Println("Future value:", futureValue)
	//fmt.Println("Future real value:", futureRealValue)
	fmt.Printf("Future value: %.2f\n", futureValue)
	fmt.Printf("Future value adjusted for inflation: %.2f\n", futureRealValue)

	// FORMAT OPTION 2:
	println("\n\nOPTION 2:")
	formattedFutureValue := fmt.Sprintf("Future value:%.2f\n", futureValue)
	formattedFutureRealValue := fmt.Sprintf("Future value adjusted for inflation:%.2f\n", futureRealValue)
	fmt.Print(formattedFutureValue, formattedFutureRealValue) // line break is not needed because Sprintf adds it automatically.

}

func outputText(text string) {
	fmt.Println(text)
}

func calculateFutureValues(investmentAmount, rateOfReturn, years float64) (futureValue float64, futureRealValue float64) {
	// Calculate future value and future real value
	futureValue = investmentAmount * math.Pow(1+rateOfReturn/100, years)
	futureRealValue = futureValue / math.Pow(1+2.5/100, years)
	return futureValue, futureRealValue
}
