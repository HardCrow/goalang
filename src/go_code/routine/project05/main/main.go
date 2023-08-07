package main

import "fmt"

func addUpper() func(n int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}
func main() {
	f := addUpper()
	fmt.Println(f(1))
}
