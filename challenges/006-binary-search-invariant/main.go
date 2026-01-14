package main

import (
	"fmt"
)

/*
invariant checks if we can eat all piles within H hours at speed k

Parameters:
  - piles: array of pile sizes
  - k: eating speed (bananas per hour)
  - H: total hours available

Returns:
  - true if all piles can be eaten within H hours at speed k
*/
func invariant(piles []int, k, H int) bool {
	var sum int = 0
	for _, v := range piles {
		sum += (v + k - 1) / k
	}

	return sum <= H
}

/*
findMaxOfArray returns the maximum value in the array

Parameters:
  - arr: array of integers

Returns:
  - maximum value found in the array
*/
func findMaxOfArray(arr []int) int {
	var maxValue int = 0
	for _, v := range arr {
		if v > maxValue {
			maxValue = v
		}
	}

	return maxValue
}

/*
minSpeed finds the minimum eating speed to finish all piles within H hours

Parameters:
  - piles: array of pile sizes
  - H: total hours available

Returns:
  - minimum speed k needed to eat all piles within H hours
*/
func minSpeed(piles []int, H int) int {
	// Binary search range: minimum speed is 1, maximum is the largest pile
	var left int = 1
	right := findMaxOfArray(piles)
	var ans int = right

	// Binary search for the minimum valid speed
	for left <= right {
		var mid int = (left + right) / 2
		// If we can finish at speed mid, try a slower speed
		if invariant(piles, mid, H) {
			ans = mid
			right = mid - 1
		} else {
			// If we can't finish, we need a faster speed
			left = mid + 1
		}
	}

	return ans
}

func main() {
	piles := []int{3, 6, 7, 11}
	var hoursOfWord int = 8

	fmt.Printf("%d\n", minSpeed(piles, hoursOfWord))
}
