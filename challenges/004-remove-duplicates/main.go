package main

import "fmt"

/*
RemoveDuplicates removes duplicate values from a sorted array in-place
and returns the count of unique elements.

After execution, the first k elements of arr will contain the unique values
in their original sorted order.

Parameters:
  - arr: sorted in non-decreasing order array

Returns:
  - k: number of unique elements
*/
func removeDuplicates(arr []int) int {
	if len(arr) == 1 {
		return 1
	}

	k := 1            // traks nextposition to write unique element
	readPointer := 0  // points to the last confirmed unique element
	writePointer := 1 // scans through array looking for new unique values

	for writePointer < len(arr) {
		if arr[readPointer] != arr[writePointer] {
			// Found a new unique element
			readPointer = writePointer
			arr[k] = arr[writePointer] // place it at position k
			k++
		}
		writePointer++
	}

	return k
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3}
	uniqueNumbers := removeDuplicates(nums)

	fmt.Printf("%d\n", uniqueNumbers)
	fmt.Printf("%v\n", nums)
}
