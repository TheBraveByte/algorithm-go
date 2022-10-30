package binary_search

import (
	"math"
)

// binarySearchRecursion : BinarySearch algorithm using recursion approach
func BinarySearchRecursion(arr []int, value, lowIndex, highIndex int) (bool, int) {
	midIndex := int(math.Floor(float64(lowIndex)+float64(highIndex)) / 2)
	if arr[midIndex] == value {
		return true, midIndex
	} else if arr[midIndex] > value {
		return BinarySearchRecursion(arr, value, lowIndex, midIndex-1)
	} else if arr[midIndex] < value {
		return BinarySearchRecursion(arr, value, midIndex+1, highIndex)
	} else {
		return false, 0
	}
}
