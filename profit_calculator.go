package main 

import (
	   "fmt"
)

func ProfitCalculator(){
	const taxRate float64 = 2.5

	var (
		revenue, expenses, profit, taxation, earningRatio, result float64
	)

	fmt.Print("kindly enter the revenue you generated this month: ")
	fmt.Scan(&revenue)

	fmt.Print("kindly lets know how much you spent in expenses: ")
	fmt.Scan(&expenses)

	profit = revenue - expenses

	taxation = (taxRate/100) * profit

	result = profit - taxation

	earningRatio = profit/result

	fmt.Println("the earning before tax is: ", profit)
	fmt.Println("the earning after tax is: ", result)
	fmt.Println("the earning ratio is: ", earningRatio)




}