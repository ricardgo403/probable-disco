package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	a := 15
	b := 999
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)
}
