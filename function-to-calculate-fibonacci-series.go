// Fibonacci sequence in Go using recursion
package main

import "fmt"

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    fmt.Println("Fibonacci Sequence:")
    for i := 0; i < 10; i++ {
        fmt.Print(fibonacci(i), " ")
    }
}

// Output
// Fibonacci Sequence:
// 0 1 1 2 3 5 8 13 21 34 
