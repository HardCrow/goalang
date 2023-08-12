package main

import "fmt"

func add(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}
func main() {
	res := add(10)
	if res != 55 {
		fmt.Println("错误")
	} else {
		fmt.Printf("成功：%v\n", res)
	}
}
