package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	//创建一个可以放三个int类型的管道
	var intChan chan interface{}
	intChan = make(chan interface{}, 4)
	fmt.Println(intChan, &intChan)
	c := Cat{"a", 1}
	intChan <- c
	intChan <- 10 //写入10进管道
	nums := 211
	intChan <- nums
	res := <-intChan
	//fmt.Println(intChan.name)//易错题   intChan此时是接口  接口里面是没有字段的
	//正确写法 要用类型断言
	c2 := res.(Cat)
	close(intChan)
	//channel关闭后可以读不可以写，close后不能加数据了
	fmt.Println(c2.Name)
	//问题解决
	fmt.Println(res)

}
