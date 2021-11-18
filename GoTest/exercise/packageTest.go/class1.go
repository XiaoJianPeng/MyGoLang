package main

import (
	"fmt"

	"./utils"
)

func OutPrint(str string) {
	fmt.Println("outPrint被调用，传入参数：", str)

	stu := utils.NewStudent(1, "张飒风", 12)
	fmt.Print(stu)

	var stu2 utils.Student
	stu2.Id = 20
	stu2.Name = "李小松"
	stu2.Age = 24
	fmt.Println("挎包引用结构体的属性必须首字母大写，", stu2)
}
