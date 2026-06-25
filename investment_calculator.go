package main 

import (
       "fmt"
        "math"
)

func InvestmentCalculator(){
	const inflationRate float64 = 2.5
	var (
		principal, years, rate, amount, inflatedAmount float64
	)
	fmt.Print("kindly enter the amount you intend to invest with prodev: ")
	fmt.Scan(&principal)

	fmt.Print("kindly enter the time you want to invest with prodev for: ")
	fmt.Scan(&years)

	fmt.Print("kindly lets know the rate they gave you at prodev for the investment: ")
	fmt.Scan(&rate)

	amount = principal * math.Pow((1 + rate/100), years)

	inflatedAmount = amount/ math.Pow((1+ inflationRate), years)

	fmt.Println("The actual amount is; ", amount)

	fmt.Println("the inflated amount is expected to be: ", inflatedAmount)



}