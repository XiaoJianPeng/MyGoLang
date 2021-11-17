package main

import (
	"fmt"
	"time"
)

var x, y int
var (
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

func main() {
	g, h := 123, "hello www"
	println(x, y, a, b, c, d, e, f, g, h)

	// var hp int
	// hp = 10
	// fmt.Println("Hello world!")
	// fmt.Println(hp)
	// fmt.Println("testVar1:")
	// testVar1()
	// fmt.Println("testVar2:")
	// testVar2()
	// 等待3秒钟
	time.Sleep(time.Second * 3)
}

func testVar1() {
	var a int = 3
	var b int = 4
	c := a + b
	fmt.Printf("a = %d, b = %d, c= %d\n ", a, b, c)
}

func testVar2() {
	var a = "Runoob"
	fmt.Println(a)

	var b int
	fmt.Println(b)

	var c bool
	fmt.Println(c)
}
