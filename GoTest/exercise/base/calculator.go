package main

import (
	"fmt"
	"reflect"
)

func add(x float32, y float32) float32 {
	return x + y
}

func subtract(x float32, y float32) float32 {
	return x - y
}

func multiply(x float32, y float32) float32 {
	return x * y
}

func divide(x float32, y float32) float32 {
	if y == 0 {
		fmt.Println("除数不能为0")
		return 0
	}
	return x / y
}

func starCalcuator() {
	var x, y float32
	var operator string
	fmt.Println("请输入第一个数字：")
	fmt.Scan(&x)
	fmt.Println(reflect.TypeOf(x).String())
	fmt.Println("请输入运算符 +，-，*，/：")
	fmt.Scan(&operator)
	fmt.Println("请输入第二个数字：")
	fmt.Scan(&y)
	var result float32 = 0
	switch operator {
	case "+":
		result = add(x, y)
	case "-":
		result = subtract(x, y)
	case "*":
		result = multiply(x, y)
	case "/":
		result = divide(x, y)
	default:
		result = add(x, y)
	}
	fmt.Println("计算结果为：", result)
	fmt.Println("按n进行下一个计算,否则退出")
	var str string
	fmt.Scan(&str)
	// fmt.Println(str)
	if str == "n" {
		starCalcuator()
	}

}

func main() {
	fmt.Println("简易计算器演示：")
	starCalcuator()
	// var typ string
	// var a float32 = 32.029
	// tyString = reflect.TypeOf(213.002)
	// typ = fmt.Sprintf("%T", a)
	// typ = reflect.TypeOf(a).String()
	// fmt.Println(typ)
}
