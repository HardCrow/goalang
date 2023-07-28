package main

import "fmt"

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err=", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println(res)

}
func main() {
	num1 := 100
	fmt.Printf("num1的类型%T,num1的值%v,num1的地址%v\n", num1, num1, &num1)
	test()
}
