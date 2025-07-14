package main

import (
	"fmt"
	"os"
)

const resultsFile = "results.txt"

func main() {
	var revenue, expenses, taxRate, ebt, profit, ratio float64
	revenue, expenses, taxRate = inputValues(revenue, expenses, taxRate)
	ebt, profit, ratio = calculateValues(revenue, expenses, taxRate)

	fmt.Println("Profit:", profit)
	fmt.Println("EBT:", ebt)
	fmt.Printf("Profit Ratio:%.3f\n", ratio)
}

func inputValues(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	//loop until user gets all values right
	for {
		fmt.Println("Enter revenue, expenses, and tax rate (seperated by spacebars):")
		fmt.Scan(&revenue, &expenses, &taxRate)
		if revenue <= 0.0 || expenses <= 0.0 || taxRate <= 0.0 {
			fmt.Print("The value of revenue,expenses or taxrate cannot be less or equal to 0!\n")
			continue
		}
		break
	}

	return revenue, expenses, taxRate
}

func calculateValues(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	fmt.Println("Calculating profit...")
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	writeResultsToFile(ebt, profit, ratio)
	return ebt, profit, ratio
}

func writeResultsToFile(ebt float64, profit float64, ratio float64) {
	resultsText := fmt.Sprintf("EBT:%.2f\nPROFIT:%.2f\nRATIO:%.2f", ebt, profit, ratio)
	os.WriteFile(resultsFile, []byte(resultsText), 0644)
	fmt.Println("Results saved to file.")
}
