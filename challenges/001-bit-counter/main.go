package main

import (
	"fmt"
)

func main() {
	var input int = 9999999999
	var hammingWeight = 0

	fmt.Printf("Input %d:\n", input)
	fmt.Print("Binary: ")

	for i := 63; i >= 0; i-- {
		if input&(1<<i) != 0 {
			fmt.Print("1")
			hammingWeight++
		} else {
			fmt.Print("0")
		}
	}

	fmt.Printf("\nCount of 1s: %d\n", hammingWeight)
}
