package main

import "fmt"

func getLastBit(n, i int, hasCarry bool) int {
	var returnValue int = 0

	if n&(1<<i) != 0 {
		returnValue = 1
	}

	if hasCarry {
		if returnValue == 0 {
			returnValue = 1
		} else {
			returnValue = 0
		}
	}

	return returnValue
}

func BinaryAdd(a, b int) int {
	for b != 0 {
		carry := a & b
		a = a ^ b
		b = carry << 1
	}

	return a
}

func main() {
	a, b := 3, 5
	result := BinaryAdd(a, b)
	fmt.Printf("Binary addition: %d + %d = %d\n", a, b, result)
	fmt.Printf("In binary: %b + %b = %b\n", a, b, result)
}
