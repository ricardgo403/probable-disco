package main

import "fmt"

func fibonacci(number int64) int64 {
	if number == 0 || number == 1 {
		return number
	}

	return (fibonacci(number-2) + fibonacci(number-1))
}

func main() {
	a := 15
	fibonacciSerie := []int64{}

	for i := 0; i < a; i++ {
		fibonacciSerie = append(fibonacciSerie, fibonacci(int64(i)))
	}
	fmt.Println(fibonacciSerie)
}
