package ps

import "fmt"

//Question:  Find triplets in array s.t. a[i] < a[j] < a[k]  and i<j<k  in linear time
/*
Solution: Use auxiliary smaller and gratter array.
smaller : smaller[i] hold the nearest smaller element index left side.
gratter : gratter[i] hold the nearest gratter element index right side
*/
func PS1() {
	var ar = []int{1, 7, 3, 5, 8, 2, 4, 9, 5, 8}
	ar = []int{12, 11, 10, 5, 6, 2, 30}
	fmt.Println(ar)
	smaller := smallar_array(ar)
	fmt.Println(smaller)
	gratter := gratter_array(ar)
	fmt.Println(gratter)
}

func smallar_array(arr []int) []int {
	size := len(arr)
	smaller := make([]int, size)
	smaller[0] = -1
	min := 0
	for i, _ := range arr {
		if arr[i] <= arr[min] {
			min = i
			smaller[i] = -1
		} else {
			smaller[i] = min
		}
	}
	return smaller
}

func gratter_array(arr []int) []int {
	size := len(arr)
	greater := make([]int, size)
	greater[size-1] = -1
	max := size - 1
	i := size - 1
	for i >= 0 {
		if arr[i] >= arr[max] {
			max = i
			greater[i] = -1
		} else {
			greater[i] = max
		}
		i--
	}
	return greater
}
