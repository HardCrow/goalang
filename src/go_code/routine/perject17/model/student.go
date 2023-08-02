package model

// 结构体名字大写 相当于public 可以在其他包内使用
type Student struct {
	Name  string
	Score int
}

// 结构体名字小写 相当于私有的 只能在本包用
type teacher struct {
	Name  string
	Score int
}

// 利用工厂模式解决
func NewTeacher(n string, s int) *teacher {
	return &teacher{
		Name:  n,
		Score: s,
	}
}
