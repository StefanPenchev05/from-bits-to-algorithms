package main

import (
	"fmt"
)

func main() {
	var x, y int = 43, 25
	k := x & y
	fmt.Printf("%d AND %d is %d\n", x, y, k)
}
