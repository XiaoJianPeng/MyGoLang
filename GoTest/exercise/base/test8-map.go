// Map集合
package main

import "fmt"

func maptest1() {
	// 创建集合
	var countryCaptialMap map[string]string
	countryCaptialMap = make(map[string]string)

	// map 插入key - value 对
	countryCaptialMap["China"] = "北京"
	countryCaptialMap["France"] = "巴黎"
	countryCaptialMap["Italy"] = "罗马"
	countryCaptialMap["Japan"] = "东京"
	countryCaptialMap["India"] = "新德里"

	for country := range countryCaptialMap {
		fmt.Println(country, "首都是：", countryCaptialMap[country])
	}

	// 查看元素在集合中是否存在
	capital, ok := countryCaptialMap["American"]
	fmt.Println(capital)
	fmt.Println(ok)
	if ok {
		fmt.Println("American 的首都是：", capital)
	} else {
		fmt.Println("American的首都不存在")
	}

	// delete 函数
	delete(countryCaptialMap, "China")
	fmt.Println("中国条目被删除；删除后的条目是：")
	for country := range countryCaptialMap {
		fmt.Println(country, "首都是：", countryCaptialMap[country])
	}

}

// 打印机空心金字塔
func printPyamid(x int) {
	for i := 0; i <= x; i++ {
		for m := 1; m <= x-i; m++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == x {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// 乘法表
func printMultTable(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}
}

func main() {
	// maptest1()
	// printPyamid(15)
	printMultTable(9)
}
