package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

// switch 用法
func switchTest() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\t", os)
	}

	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("tomorrow")
	case today + 2:
		fmt.Println("in two days")
	default:
		fmt.Println("Too far away")
	}
}

// defer 用法, defer 会把数据放入到栈中， 先入后出原则
func deferTest() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")

}

func main() {
	// switchTest()
	// deferTest()
	var tyString string
	tyString = reflect.TypeOf(213.002)
	fmt.Println(tyString)
}
