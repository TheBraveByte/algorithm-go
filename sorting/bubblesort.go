package sorting

import "fmt"

//BubbleSort : this algorithms 
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
			fmt.Println(arr)
		}
	}
	return arr
}
