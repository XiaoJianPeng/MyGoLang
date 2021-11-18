package main

import (
	"fmt"

	"./utils"
)

func OutPrint(str string) {
	fmt.Println("outPrint被调用，传入参数：", str)

	stu := utils.NewStudent(1, "张飒风", 12)
	fmt.Print(stu)
}
