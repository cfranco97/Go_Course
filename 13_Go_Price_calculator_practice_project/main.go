package main

import (
	"example.com/price_calculator/prices"
)

func main() {

	taxRates := []float64{0, 0.5}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}
}
