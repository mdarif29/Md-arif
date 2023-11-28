package main

import (
	"fmt"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func main() {
	var num1, num2 int

	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)

	if num1 <= 0 || num2 <= 0 {
		fmt.Println("Please enter positive integers.")
	} else {
		result := lcm(num1, num2)
		fmt.Printf("LCM of %d and %d is %d\n", num1, num2, result)
	}
}
