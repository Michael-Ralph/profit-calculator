package main

import (
	"fmt"
)

func main() {
	// Variables
	var revenue, expenses, taxRate float64

	// Get variables
	fmt.Print("What is your total revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("What is your total expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("What is your tax rate: ")
	fmt.Scan(&taxRate)

	// Calculations
	// EBT (Earnings Before Tax) = Revenue - Operating Expenses - Interest Expenses

	// Tax Amount = EBT × Tax Rate

	// Net Profit = Revenue - Operating Expenses - Interest Expenses - (EBT × Tax Rate)

	// Pretax Margin Ratio = EBT / Revenue * 100

	// Outputs

}
