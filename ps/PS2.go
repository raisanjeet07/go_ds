package ps

import "fmt"

//Question:  Given an array, after removing an element, the sum of odd and even indexed elements are equal.
/*
ex1: 2,1,6,4
	After removing 1,
	2 + 4 = 6
ex2: 2,6,6
	After removing 2,
	6 = 6
ex2: 2,2,6
	After removing 6,
	2 = 2
*/

func PS2() {
	arr := []int{2, 2, 1, 2, 2}
	fmt.Println(arr)
	index := 0
	count := 0
	if len(arr) < 3 {
		fmt.Println(0)
		return
	}
	odd_even_sum := get_even_odd_sum_dp(arr)
	// while removing the element at any index, in the right side of the elements, odd indexd elements will become even and vice-versa
	first_half_even_sum := 0
	first_half_odd_sum := 0
	f_odd_sum := odd_even_sum[len(arr)-1]
	f_even_sum := odd_even_sum[len(arr)-2]
	if len(arr)%2 != 0 {
		f_odd_sum = odd_even_sum[len(arr)-2]
		f_even_sum = odd_even_sum[len(arr)-1]
	}
	for index < len(arr) {
		tmp_f_even_sum := f_even_sum // 8
		tmp_f_odd_sum := f_odd_sum   // 5
		if index%2 == 0 {            // removing even one
			first_half_even_sum = odd_even_sum[index] - arr[index]
			if index >= 1 {
				first_half_odd_sum = odd_even_sum[index-1]
			} else {
				first_half_odd_sum = 0
			}
			tmp_f_even_sum = f_even_sum - arr[index]
		} else { // removing the odd one
			first_half_odd_sum = odd_even_sum[index] - arr[index] // 0
			if index >= 1 {
				first_half_even_sum = odd_even_sum[index-1]
			} else {
				first_half_even_sum = 0
			}
			tmp_f_odd_sum = f_odd_sum - arr[index]
		}
		total_odd_sum := first_half_odd_sum + (tmp_f_even_sum - first_half_even_sum)
		total_even_sum := first_half_even_sum + (tmp_f_odd_sum - first_half_odd_sum)
		if total_odd_sum == total_even_sum {
			count++
		}
		index++
	}
	fmt.Println(count)
}

// this method will called if the passed array has size >2
func get_even_odd_sum_dp(arr []int) []int {
	size := len(arr)
	dp := make([]int, size)
	dp[0] = arr[0]
	dp[1] = arr[1]
	index := 2
	for index < size {
		dp[index] = dp[index-2] + arr[index]
		index++
	}
	return dp
}
