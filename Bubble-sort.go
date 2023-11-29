This Go program uses the bubble sort algorithm to sort an array.
The bubbleSort function iterates through the array multiple times, each time swapping adjacent elements if they are in the wrong order. 
This process continues until the entire array is sorted. The main function creates an array and then calls the bubbleSort function to sort it. 
Finally, it prints the sorted array.

WAP OF BUBBLE SORT
package main

import "fmt"

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Unsorted array:", arr)
	bubbleSort(arr)
	fmt.Println("Sorted array:", arr)
}
