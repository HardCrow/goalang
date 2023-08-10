package service

import (
	"go_code/project/XXXLproject/entry"
)

type CustomerService struct {
	customers    []entry.Customer
	CustomerNums int
}

func NewCustomerService() *CustomerService {
	//初始化用户
	customerService := &CustomerService{} //定义一个 Customer结构的对象为 customerService
	customerService.CustomerNums = 1
	customer := entry.Customer{
		1,
		"huang",
		"nan",
		12,
		"21545102",
		"shdaudahd",
	}
	customerService.customers = append(customerService.customers, customer) //在costomers的切片里面加了一个 costomer对象
	return customerService                                                  //返回一个Customer结构类型的对象给函数
}
func (this *CustomerService) List() []entry.Customer {
	return this.customers //记录并保存客户记录
}
func (this *CustomerService) Add(customer entry.Customer) bool {
	this.CustomerNums += 1
	customer.Id = this.CustomerNums
	this.customers = append(this.customers, customer)
	return true
}
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}
	this.customers = append(this.customers[:index], this.customers[index+1:]...) //切片从零开始取到index到但是不含inde x
	return true
}
func (this *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			index = i
		}
	}
	return index
}
