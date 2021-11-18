package main

import "fmt"

func main() {
	fmt.Println("package包引入练习")
	OutPrint("Hellow  world!")

	var p Person
	p.name = "王小二"
	p.age =34
	p.gender = "男"
	fmt.Println(p)
}
