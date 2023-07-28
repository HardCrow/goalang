package main

import "fmt"

func sum(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[0]
	}
	return sum
}
func main() {
	res := sum(10, 1, 20, 33, 5, -54)
	fmt.Println(res)
}
