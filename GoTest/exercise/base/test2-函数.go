package main

import "fmt"

func main() {
	// const (
	// 	a = iota //0
	// 	b        //1
	// 	c        //2
	// 	d = "ha" //独立值，iota += 1
	// 	e        //"ha"   iota += 1
	// 	f = 100  //iota +=1
	// 	g        //100  iota +=1
	// 	h = iota //7,恢复计数
	// 	i        //8
	// )
	// fmt.Println(a, b, c, d, e, f, g, h, i)
	// forTest()

	// a, b := swap("hello ", "world")
	// fmt.Println(a, b)

	// 测试闭包函数
	nextnumber := getSequence()
	for i := 1; i < 5; i++ {
		fmt.Println(nextnumber())
	}
}

func forTest() {
	for i := 0; i < 10; i++ {
		fmt.Println("这是一个循环", i)
	}
}

func swap(x, y string) (string, string) {
	return y, x
}

// 闭包函数
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
