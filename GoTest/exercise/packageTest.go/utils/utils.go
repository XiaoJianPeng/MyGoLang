package utils

import "fmt"

type Student struct {
	Id   int
	Name string
	Age  int
}

func NewStudent(id int, name string, age int) Student {
	fmt.Println("utils.NewStudent被调用")
	return Student{
		Id:   id,
		Name: name,
		Age:  age,
	}
}
