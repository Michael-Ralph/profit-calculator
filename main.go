package main

import (
	"fmt"
	"os"
)

func main() {
	// Variables
	var revenue, expenses, taxRate float64

	// Get variables
	revenue, expenses, taxRate = getVariables()

	// Calculations
	ebt, _, netProfit, pmr := calcs(revenue, expenses, taxRate)

	// Write data to file
	writeData(ebt, netProfit, pmr)

	// Outputs
	fmt.Printf("Earnings Before Tax: R%.2f\n", ebt)
	fmt.Printf("Net profit: R%.2f\n", netProfit)
	fmt.Printf("Pretax Margin Ratio: %.2f\n", pmr)
}

func getVariables() (float64, float64, float64) {
	var revenue, expenses, taxRate float64
	fmt.Print("What is your total revenue: R")
	fmt.Scan(&revenue)
	if revenue <= 0 {
		fmt.Println("Revenue can not be less than R0.")
		os.Exit(1)
	}
	fmt.Print("What are your total expenses: R")
	fmt.Scan(&expenses)
	if expenses <= 0 {
		fmt.Println("Expenses can not be less than R0.")
		os.Exit(1)
	}
	fmt.Print("What is your countries tax rate(%): ")
	fmt.Scan(&taxRate)
	if taxRate <= 0 {
		fmt.Println("tax rate can not be less than 0%.")
		os.Exit(1)
	}

	return revenue, expenses, taxRate
}

func calcEBT(revenue, expenses float64) float64 {
	EBT := revenue - expenses
	return EBT
}

func taxCalc(ebt, taxRate float64) float64 {
	taxAmount := ebt * (taxRate / 100)
	return taxAmount
}

func netProfit(ebt, taxAmount float64) float64 {
	netProfit := ebt - (taxAmount)
	return netProfit
}

func calcPMR(ebt, revenue float64) float64 {
	pmr := (ebt / revenue) * 100
	return pmr
}

func calcs(revenue, expenses, taxRate float64) (float64, float64, float64, float64) {
	ebt := calcEBT(revenue, expenses)
	taxAmount := taxCalc(ebt, taxRate)
	netProfit := netProfit(ebt, taxAmount)
	pmr := calcPMR(ebt, revenue)

	return ebt, taxAmount, netProfit, pmr
}

func writeData(ebt, netProfit, pmr float64) {
	dataStr := fmt.Sprintf("EBT : R%.2f, Net profit : R%.2f, PMR : %.2f", ebt, netProfit, pmr)
	os.WriteFile("results.txt", []byte(dataStr), 0644)
}
