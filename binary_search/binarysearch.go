package binary_search

import (
	"log"
	"math"

)

func BinarySearch(arr []int, checkValue int) int {

	var midIndexInt int
	count_search := 0

	lowIndex := 0
	highIndex := len(arr) - 1
	for lowIndex <= highIndex {
		midIndex := math.Round((float64(lowIndex) + float64(highIndex)) / 2)
		midIndexInt = int(midIndex)
		midValue := arr[midIndexInt]
		if midValue == checkValue {
			log.Println("count for the search is:", count_search)
			return midIndexInt
		}
		if midValue < checkValue {
			count_search += 1
			lowIndex += 1
		} else {
			highIndex -= 1
			count_search += 1
		}

	}
	log.Println("No result :: invalid search")
	return 0
}
