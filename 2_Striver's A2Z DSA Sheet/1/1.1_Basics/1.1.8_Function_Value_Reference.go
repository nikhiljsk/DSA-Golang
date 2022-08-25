package main

import (
	"fmt"
)

func ref(a *int) *int {
	*a += 1
	return a
}

func value(a int) int {
	a += 1
	return a
}

func main() {
	var a int
	fmt.Println("Before value:", a)
	value(a)
	fmt.Println("After value:", a)
	fmt.Println("Before ref:", a)
	ref(&a)
	fmt.Println("After ref:", a)
}
