package main

import "fmt"

func Conditionals(){

	fmt.Println("Welcome to the GO Bank !")
	fmt.Println("What do you want to do? ")
	fmt.Println("1. Check Balance")
	fmt.Println("2- Deposit money")
	fmt.Println("3- withdrawal")
	fmt.Println("4- Exit")
   

	var accountBalance float64
	accountBalance = 50000
	var choice int
	fmt.Print("kindly enter the your choice : ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your account balance is : ", accountBalance)
	}else if choice == 2{
		fmt.Println("Your chose deposit: ")
		var deposit float64 
		// newBalance float64
		fmt.Println("kindly input the amount you want to deposit: ")
		fmt.Scan(&deposit)
		// newBalance = deposit + accountBalance
		fmt.Println("Your you deposited ", deposit)
		accountBalance += deposit
		fmt.Println("your new account balance is :", accountBalance)
	}else if choice == 3 {
		fmt.Print("You chose withdrawal ")
		fmt.Print("how much do you want to withdraw: ")
		var withdrawal float64
		fmt.Scan(&withdrawal)
		fmt.Println("Please take your cash: ", withdrawal)
		accountBalance -= withdrawal
		fmt.Println("your new balance is : ", accountBalance)
	}else {
		fmt.Println(" exiting ......")
		
	}

 fmt.Println("thank you")
} 