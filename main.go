package main

import (
	"fmt"
)

func main() {
	// Variables
	var revenue, expenses, taxRate float64

	// Get variables
	fmt.Print("What is your total revenue: R")
	fmt.Scan(&revenue)
	fmt.Print("What is your total expenses: R")
	fmt.Scan(&expenses)
	fmt.Print("What is your tax rate: ")
	fmt.Scan(&taxRate)

	// Calculations
	// EBT (Earnings Before Tax) = Revenue - Operating Expenses - Interest Expenses
	EBT := revenue - expenses
	// Tax Amount = EBT × Tax Rate
	taxAmount := EBT * (taxRate / 100)
	// Net Profit = Revenue - Operating Expenses - Interest Expenses - (EBT × Tax Rate)
	netProfit := EBT - (taxAmount)
	// Pretax Margin Ratio = EBT / Revenue * 100
	PMR := (EBT / revenue) * 100

	// Outputs
	fmt.Printf("Earnings Before Tax: R%.2f\n", EBT)
	fmt.Printf("Net profit: R%.2f\n", netProfit)
	fmt.Printf("Pretax Margin Ratio: %.2f\n", PMR)
}
