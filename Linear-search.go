// Algo: Approach
def linear_search(arr, x):
    """
    Perform a linear search on the input array 'arr' to find the element 'x'.
    Returns the index of the element if found, else -1.
    """
    n = len(arr)
    # Best case: x is at the beginning of the array
    if arr[0] == x:
        return 0
    # Worst case: x is at the end of the array
    elif arr[n-1] == x:
        return n-1
    # Average case: x is somewhere in the middle of the array
    else:
        low = 0
        high = n-1
        while low <= high:
            mid = (low + high) // 2
            if arr[mid] == x:
                return mid
            elif arr[mid] < x:
                low = mid + 1
            else:
                high = mid - 1
        return -1

// Code: need to make it work
// WAP OF LINEAR SEARCH
package main

import (
   "fmt"
)

func linearSearch(arr []int, x int) int {
   for i := 0; i < len(arr); i++ {
   	if arr[i] == x {
   		return i
   	}
   }
   return -1
}

func main() {
   arr := []int{2, 3, 4, 10, 40}
   x := 10

   result := linearSearch(arr, x)
   if result != -1 {
   	fmt.Printf("Element is present at index %d", result)
   } else {
   	fmt.Printf("Element is not present in array")
   }
}
