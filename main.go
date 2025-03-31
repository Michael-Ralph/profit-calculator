package main

import (
	"fmt"
)

func main() {
	// Variables
	var revenue, expenses, taxRate float64

	// Get variables
	revenue, expenses, taxRate = getVariables()

	// Calculations
	// EBT (Earnings Before Tax) = Revenue - Operating Expenses - Interest Expenses
	ebt := calcEBT(revenue, expenses)
	// Tax Amount = EBT × Tax Rate
	taxAmount := taxCalc(ebt, taxRate)
	// Net Profit = Revenue - Operating Expenses - Interest Expenses - (EBT × Tax Rate)
	netProfit := netProfit(ebt, taxAmount)
	// Pretax Margin Ratio = EBT / Revenue * 100
	pmr := calcPMR(ebt, revenue)

	// Outputs
	fmt.Printf("Earnings Before Tax: R%.2f\n", ebt)
	fmt.Printf("Net profit: R%.2f\n", netProfit)
	fmt.Printf("Pretax Margin Ratio: %.2f\n", pmr)
}

func getVariables() (float64, float64, float64) {
	var revenue, expenses, taxRate float64
	fmt.Print("What is your total revenue: R")
	fmt.Scan(&revenue)
	fmt.Print("What are your total expenses: R")
	fmt.Scan(&expenses)
	fmt.Print("What is your countries tax rate(%): ")
	fmt.Scan(&taxRate)

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
