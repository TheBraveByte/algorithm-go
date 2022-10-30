package binary_search

import (
	"math"
)
// TwoBalls : this algorithm is an optimized binary search based on the 
// condition or scenerio. using the Big notation of O(N^0.5)
// Question: Given two crstal ball that will break if dropped from high enough
// distance .Determine the exact spot or position in which it will break
// in the optimized way
func TwoBalls(arr []bool) int {
	jump_search := int(math.Floor(math.Sqrt(float64(len(arr)))))
	i := jump_search

	// finding the breakspot for the first ball
	for ; i < len(arr); i += jump_search {
		if arr[i] {
			break
		}
	}

	// finding the break spot for the second ball
	i -= jump_search
	for j := 0; j < jump_search && i < len(arr); j, i = j+1, i+1 {
		if arr[i] {
			return i
		}
		
	}
	return -1
}
