package main

import (
	"fmt"
)

func main() {
	var intChan chan int
	// 给管道写入数据不能超过它的容量
	intChan = make(chan int, 3)
	fmt.Printf("intChan的值=%v, intChan的地址%v \n", intChan, &intChan)

	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 29

	fmt.Printf("channel len =%v cap=%v  value=%v \n", len(intChan), cap(intChan), intChan)

	// 超出chan长度会报错 all groutines are asleep - deadlock
	for i := 0; i < 3; i++ {
		fmt.Printf("第%d次取出的值=%d, intChan的长度%v \n", i, <-intChan, len(intChan))
	}
}
