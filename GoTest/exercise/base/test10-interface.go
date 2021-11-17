// Go语言类型转换 + 接口

package main

import "fmt"

func test1() {
	var sum int = 15
	var count int = 4
	var mean float32
	mean = float32(sum) / float32(count)
	fmt.Printf("mean的值是：%f\n", mean)
}

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

func (nokiaPhone NokiaPhone) say() {
	fmt.Println("调用了say函数")
}

type Iphone struct {
}

func (ipone Iphone) call() {
	fmt.Println("I am Ipone, I can call you!")
}

func main() {
	// test1()
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(Iphone)
	phone.call()

	var nokia NokiaPhone
	nokia.say()
	nokia.call()
}
