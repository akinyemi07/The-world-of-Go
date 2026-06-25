package main

import "fmt"
import "math"

func InvestmentCalcualtor () {
var (
	principal, rate, years float64 = 100000, 5.5, 10
)
 var future = principal * math.Pow((1 + rate / 100), years)
 fmt.Println(future)
 
}