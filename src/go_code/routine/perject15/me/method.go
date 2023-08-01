package main

import "fmt"

type Person struct {
	Name string
}

// 要求一： 给person结构添加一个方法speak输出xxx是好人
func (p Person) test() { //注意这个和java不一样   （p person）类似于判断对象是否是person类
	fmt.Println(p.Name) //这个方法是和p person绑定的 不和其他绑定 如下dag对象不能调用该方法
}
func (person Person) speak() {
	fmt.Println(person.Name + "是一个好人")
}
func (person2 Person) jisuan() {
	res := 0
	for i := 0; i <= 1000; i++ {
		res += i
	}
	fmt.Println(res)
}
func (p Person) jisuan2(n int) {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	fmt.Println(res)
}
func (p Person) getSum(n1 int, n2 int) int {
	return n1 + n2
}

type Dog struct {
}

func main() {
	var p Person
	p.Name = "tom"
	p.test()
	p.speak()
	p.jisuan()
	p.jisuan2(10)
	res := p.getSum(10, 20)
	fmt.Println(res)
	//  var dog Dog    //测试dog对象是否可以调用test（）方法
	//  dog.test()     //编译时异常 test（）未绑定该对象
}
