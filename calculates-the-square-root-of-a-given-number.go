package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64

	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if num < 0 {
		fmt.Println("Cannot calculate the square root of a negative number.")
	} else {
		result := math.Sqrt(num)
		//2f" tells the printf method to print a floating point value (the double, x, in this case) with 2 decimal places. 
		fmt.Printf("Square root of %.2f is %.2f\n", num, result)
	}
}
