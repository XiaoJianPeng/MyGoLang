package main

import (
	"fmt"
)

// 向管道写入数据
func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Println("Wrire data ", i)
	}
	// 写入完成关闭管道
	close(intChan)
}

// 从管道读取数据
func redData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("read data 读到数据%v \n", v)
	}
	//读取完成后给管道exitChan写入true的标记，并关闭exitChan
	exitChan <- true
	close(exitChan)
}

func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool)
	go writeData(intChan)
	go redData(intChan, exitChan)
	// 循环判断是否完成任务,完成后则退出循环
	for {
		_, ok := <-exitChan
		if ok {
			fmt.Println("任务完成")
			break
		}
	}

}
