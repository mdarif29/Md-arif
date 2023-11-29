This Go program uses the QuickSort algorithm to sort an array of random integers. 
It creates an array of the specified size, fills it with random integers, and then sorts the array using the QuickSort algorithm.

The time it takes to sort the array is measured and displayed along with the sorted array.
WAP OF QUICK SORT
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	size := 100
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Intn(1000)
	}

	start := time.Now()
	quickSort(data)
	elapsed := time.Since(start)

	fmt.Println("Sorted: ", data)
	fmt.Println("Time: ", elapsed)
}

func quickSort(data []int) {
	if len(data) <= 1 {
		return
	}

	pivot := data[len(data)/2]
	var less []int
	var greater []int

	for _, item := range data {
		if item < pivot {
			less = append(less, item)
		} else {
			greater = append(greater, item)
		}
	}

	quickSort(less)
	quickSort(greater)

	i := 0
	for _, item := range less {
		data[i] = item
		i++
	}

	data[i] = pivot
	i++

	for _, item := range greater {
		data[i] = item
		i++
	}
}
