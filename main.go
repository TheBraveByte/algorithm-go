package main

import (
	"fmt"
	"log"

	"github.com/yusuf/algorithm-go/binary_search"
	"github.com/yusuf/algorithm-go/hanoitower"
	"github.com/yusuf/algorithm-go/maxnumber"
	"github.com/yusuf/algorithm-go/sorting"
)

func main() {
	// Recursion & BackTracking
	fmt.Println("Starting Tower of Hanio algorithms")
	hanoitower.TowerOfHanio(3, "A", "B", "C")

	log.Println("max number in an array of integers")
	maxnumber.MAX()

	// The binary search algorithm example
	log.Println("Binary search algorithms")
	fmt.Println(binary_search.BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8))

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Println(binary_search.BinarySearchRecursion(arr, 11, 0, len(arr)-1))
	log.Println("Two ball binary search")
	fmt.Println(binary_search.TwoBalls([]bool{false, false, false, false, false, false, true, true, true}))

	fmt.Println()
	log.Println(sorting.BubbleSort([]int{1, 2, 3, 5, 4, 7, 6, 9, 8}))
}
