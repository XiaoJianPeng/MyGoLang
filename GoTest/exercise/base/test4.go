/**
* 指针
**/
package main

import "fmt"

func test1() {
	var a int = 10
	fmt.Printf("变量地址：%x\n", &a)
}

func test2() {
	var a int = 20
	var ip *int
	ip = &a
	fmt.Printf("a的变量地址是：%x\n", &a)
	fmt.Printf("ip的变量地址是：%x\n", ip)
	fmt.Printf("ip的值是：%d\n", *ip)
}

func test3() {
	var ptr *int
	if ptr == nil {
		fmt.Printf("ptr 的值为nil\n")
	}
	fmt.Printf("ptr 的值为 : %x\n", ptr)
}

func main() {
	// test1()
	// test2()
	test3()
}
