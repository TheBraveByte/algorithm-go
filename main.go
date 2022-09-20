package main

import "fmt"

func main() {

	// The binary search algorithm example
	fmt.Println(BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8))
	
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,11,12,13,14,15}
	fmt.Println(binarySearchRecursion(arr, 11, 0, len(arr)-1))
}
