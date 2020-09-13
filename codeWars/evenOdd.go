package main

import (
	"fmt"
)

//Problem: Accept a number and determine if it is Even or Odd

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func main() {
	fmt.Println("Enter a number: ")
	var inputNum int
	fmt.Scanln(&inputNum)
	fmt.Println("The number is", EvenOrOdd(inputNum))
}
