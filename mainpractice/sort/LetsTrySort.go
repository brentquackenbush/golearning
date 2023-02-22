package main

import (
	"fmt"
)

func main() {
	xi := []int{6, 2, 1, 5, 3, 4}
	fmt.Println(xi)
	//sort.Ints(xi)
	fmt.Println(sortInts(xi))
}

/**
 * Bubble sort
 * O(n^2) time compelxity bleh
 */
func sortInts(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	return arr
}
