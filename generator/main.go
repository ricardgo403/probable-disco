package main

import "fmt"

func oddGenerator() func() uint {
	i := uint(1)
	return func() uint {
		var odd = i
		i += 2
		return odd
	}
}

func main() {
	nextOdd := oddGenerator()
	for i := 0; i < 10; i++ {
		fmt.Println(nextOdd())
	}
}
