package main

import "fmt"

// go run main.go Person.go class1.go
func main() {
	fmt.Println("package包引入练习")
	// 同目录引用同一个包下不同文件里的内容，必须同时执行run  `go run main.go Person.go class1.go`
	OutPrint("Hellow  world!")

	var p Person
	p.name = "王小二"
	p.age = 34
	p.gender = "男"
	fmt.Println(p)
}
