package main

import "fmt"

func main() {
	fmt.Println("Sum:", Add(3, 5))
}

func Add(a int, b int) int {
	return a + b
}