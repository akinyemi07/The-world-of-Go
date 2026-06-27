package main

import (
	"fmt"
)

func Temp() {
	// a program that prints the range of temperatures entered by the users

	const lower, step float64 = 0.0, 20.0
	var (
		upper, celcius, fahr float64
	)
	fmt.Print("kindly enter the value of the upper temperature marker: ")
	fmt.Scan(&upper)

	for fahr = lower; fahr <= upper; fahr += step {
		celcius = (5.0 / 9.0) * (fahr - 32)
		fmt.Printf("%3.2fF : %6.2fC \n", fahr, celcius)
	}
	/*package main

	import "fmt"

	func Temp() {
		const lower float64 = 0

		var (
			upper   float64
			fahr    float64
			celsius float64
		)

		fmt.Print("Enter the upper temperature: ")
		fmt.Scan(&upper)

		for fahr = lower; fahr <= upper; fahr++ {
			celsius = (5.0 / 9.0) * (fahr - 32)
			fmt.Printf("%6.2f°F = %6.2f°C\n", fahr, celsius)
		}
	}*/
}
