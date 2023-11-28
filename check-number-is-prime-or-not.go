package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	
	maxDivisor := int(math.Sqrt(float64(n)))

	for i := 2; i <= maxDivisor; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	if isPrime(number) {
		fmt.Printf("%d is a prime number.\n", number)
	} else {
		fmt.Printf("%d is not a prime number.\n", number)
	}
}
