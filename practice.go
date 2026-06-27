package main

import (
	"fmt"
	"math"
)

func practice() {

	//a program that has to do with print formatting and recieving input

	//lets try the area of a circle

	const pi float64 = 3.142
	var (
		area, radius float64
	)
	fmt.Print("kindly enter the radius of the circle we are trying to calculate: ")
	fmt.Scan(&radius)
	area = pi * math.Pow(radius, 2)

	fmt.Println("the area of the circle is: ", area)

}
