package main

import "fmt"

/*
 defer 语句会将函数推迟到外层函数返回之后执行。
推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
多层函数调用时， 从最外层函数一次执行到最里层函数
*/
func main() {
	defer print()
	// 第二行输出
	defer fmt.Println("world")

	// 第一行输出
	fmt.Println("hello")
}

func print() {
	// 第四行输出
	defer fmt.Println("第二行")
	// 第三行输出
	fmt.Println("第一行")

}
