package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNums", cpuNum)
	// go1.8以后默认使用最大CPU数
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}
