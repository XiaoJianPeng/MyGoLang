/*
* 1。 主线程是一个物理线程，直接作用在cpu上的。是重量级的，非常耗CPU资源
* 2.协程从主线程开启的，是轻量级的线程，是逻辑态，对资源消耗相对小
* 3. Golang 的协程机制是重要特点，可以轻松开启上万个协程。其他语言的并发机制是一般基于线程的，开启过多的线程，资源耗费大，
* 这里就凸显Golang在并发上的优势了

goroutine的调度模型 MPG
*/
package main

import (
	"fmt"
	"strconv"
	"time"
)

func func1() {
	for i := 0; i < 10; i++ {
		fmt.Println("func1() hello world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func func2() {
	for i := 0; i < 10; i++ {
		fmt.Println("222===22")
		time.Sleep(time.Second)
	}
}

//主线程
func main() {
	// 开启一个协程，如果主线程执行完了，协程没有执行完，也会退出程序
	go func1()
	// 开启一个协程
	go func2()

	for i := 0; i < 10; i++ {
		fmt.Println("main() hello world =====" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
