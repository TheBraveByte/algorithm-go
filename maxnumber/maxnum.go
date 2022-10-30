package maxnumber

import "fmt"

func MAX() {
	arr := []int{653, 8, 1, 2, 4, 0, 9, 10, 15, 22, 99}
	max := 0
	for x := range arr {
		check := arr[x]
		if check > max {
			max = check
		} else {
			check = max

		}
	}
	fmt.Println(max)

}
