package main

import "fmt"

func sum(args ...int) int {
	masGrande := 0
	for _, v := range args {
		if v > masGrande {
			masGrande = v
		}
	}
	return masGrande
}

func main() {
	a := []int{987, 564, 1000, 999, 1, 2, 0, -1, 1001, 9845789}
	fmt.Println(sum(a...))
}
