package main

import (
	"fmt"
	"go_code/project/XXXLproject/entry"
	"go_code/project/XXXLproject/service"
)

type customerView struct {
	k               string
	loop            bool
	CustomerService *service.CustomerService
}

func (this *customerView) list() {
	customers := this.CustomerService.List() //客户初始化 有一个自己提前写好的用户
	fmt.Println("---------客户列表---------")
	fmt.Println("编号\t姓名\t\t\t性别\t\t年龄\t\t电话\t\t\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("--------客户列表完成-------")
}
func (this *customerView) mainMenu() {
	for {
		fmt.Println("--------------客户管理系统---------------")
		fmt.Println("1.添加客户")
		fmt.Println("2.修改客户")
		fmt.Println("3.删除客户")
		fmt.Println("4.客户列表")
		fmt.Println("5.退   出")
		fmt.Println("请选择1-5:")
		fmt.Scanln(&this.k)
		switch this.k {
		case "1":
			this.Add()
		case "2":
			fmt.Println()
		case "3":
			this.Delete()
		case "4":
			this.list()
		case "5":
			this.loop = false
		default:
			fmt.Println("你的输入有误")
		}
		if !this.loop { //loop 为假则跳出
			break
		}
	}
	fmt.Println("你退出了用户管理系统")
}
func (this *customerView) Delete() {
	fmt.Println("--------------删除用户------------------")
	fmt.Println("请选择删除的id（-1退出）:")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("是否要删除y/n:")
	choise := ""
	fmt.Scanln(&choise)
	if choise == "y" {
		if this.CustomerService.Delete(id) {
			fmt.Println("删除成功")
		} else {
			fmt.Println("删除失败,请重新输入")
			this.Delete()
		}
	}
}

func (this *customerView) Add() {
	fmt.Println("添加客户")
	fmt.Println("姓名") //id是系统分配
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话号码")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电子邮件")
	email := ""
	fmt.Scanln(&email)
	customer2 := entry.NewCustomer2(name, gender, age, phone, email)
	if this.CustomerService.Add(customer2) {
		fmt.Println("添加成功")
	}
}
func main() {
	var test *service.CustomerService
	customerView := customerView{
		"",
		true,
		test,
	}
	customerView.CustomerService = service.NewCustomerService()
	customerView.mainMenu()

}
