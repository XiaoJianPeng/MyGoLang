package utils

import "fmt"

type Student struct {
	id   int
	name string
	age  int
}

func NewStudent(id int, name string, age int) Student {
	fmt.Println("utils.NewStudent被调用")
	return Student{
		id:   id,
		name: name,
		age:  age,
	}
}
