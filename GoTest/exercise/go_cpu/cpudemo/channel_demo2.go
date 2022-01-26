package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	var allChan chan interface{}
	allChan = make(chan interface{}, 3)

	allChan <- 10
	allChan <- "tomeweq"
	allChan <- Cat{"花猫", 2}

	<-allChan
	<-allChan

	newCat := <-allChan
	fmt.Printf("newCat = %T, newCat = %v \n", newCat, newCat)
	//使用断言类型转化
	cat := newCat.(Cat)

	fmt.Printf("newCat.Name =%v \n", cat.Name)

	//管道遍历 与关闭
	//新建管道，并写入数据
	intChan2 :=make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan2<- i*2
	}
	// 遍历管道时，如果channel没有关闭，则会出现deadlock错误
	// 遍历管道时，如果channel已经关闭，则遍历完成后，会退出遍历
	close(intChan2)
	for v :=range intChan2 {
		fmt.Println("v=", v)
	}
}
