package main

import (
	"bytes"
	"fmt"
	"math/rand"
)

func main() {
	// The binary search algorithm example
	fmt.Println(BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8))

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Println(binarySearchRecursion(arr, 11, 0, len(arr)-1))

	// Recursion & BackTracking
	fmt.Println("Starting Tower of Hanio algorithms")
	TowerOfHanio(3, "A", "B", "C")

	value := []byte("1011001")
	//var r *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	for x := 0; x < len(value); x++ {
		rand.Shuffle(len(value), func(i, j int) {
			value[i], value[j] = value[j], value[i]
		})

		for y, _ := range value {
			if y < 4 {
				ok := bytes.EqualFold(value[y:y+3], []byte("101"))
				if !ok {
					fmt.Printf("\n%c", value)
					break
				}
			}
		}

	}
	MAX()
}
