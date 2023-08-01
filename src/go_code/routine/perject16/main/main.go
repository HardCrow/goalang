package main

import "fmt"

// 关于方法和函数值传递和引用传递的问题，与主函数怎么写无关（也就是说与调用者无关） 看方法或者函数接受的类型是否传递地址还是值
type Person struct {
	Name string
}

func (p Person) test() { //方法person类型加了*说明是引用传递 没有加*号就是值传递    和主函数怎么写无关
	p.Name = "jjjj"
	fmt.Println("name", p.Name)
}
func test02(p Person) { //func test02(p *Person)就成了引用传递 和主函数里怎么写无关
	p.Name = "ccc"
	fmt.Println(p.Name + "函数内部输出")
}
func main() {
	p := Person{"tom"}  //方法的调用
	p.test()            //强调方法的调用是值传递
	fmt.Println(p.Name) //强调方法的调用是值传递
	(&p).test()
	fmt.Println(p.Name) //即使用取地址的方式去调用方法也是值传递
	fmt.Println("...............................................")
	test02(p)
	fmt.Println(p.Name + "主函数输出") //也是值传递
	// test02(&p)
	// fmt.Println(p.Name)
}
