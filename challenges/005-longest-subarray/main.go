package main

import "fmt"

/*
FindMaxSubArray finds the maximum length of contiguous subarray whose sum is <= S

Parameters:
  - arr: unsorted array with positive values
  - maxSum: the sum that the subarray should be less or equal
*/
func findMaxSubArray(arr []int64, maxSum int64) int {
	if len(arr) == 0 {
		return 0
	}

	l := 0
	currLength := 0
	maxLength := 0
	var currSum int64 = 0

	for r := 0; r < len(arr); r++ {
		currSum += arr[r]
		for currSum > maxSum && l <= r {
			currSum -= arr[l]
			l++
		}

		currLength = r - l + 1
		if currLength > maxLength {
			maxLength = currLength
		}
	}

	return maxLength
}

func main() {
	nums := []int64{2, 1, 5, 1, 3, 2}
	var s int64 = 7

	fmt.Println(findMaxSubArray(nums, s))
}
