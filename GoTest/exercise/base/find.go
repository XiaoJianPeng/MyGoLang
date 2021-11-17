// 查找
package main

import "fmt"

var names = []string{"张三", "李四", "王二", "李三", "赵六"}
var arrint = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 顺序查找方式1
func find1(str string) {
	for i, name := range names {
		if str == name {
			fmt.Printf("找到了%v,下标%v\n", name, i)
			break
		} else if i == len(names)-1 {
			fmt.Println("查不到你要找的人名")
		}
	}
}

//顺序查找方式2
func find2(str string) {
	index := -1
	for i, name := range names {
		if str == name {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Printf("找到了%v,下标%v\n", names[index], index)
	} else {
		fmt.Println("查不到你要找的人名")
	}
}

// 二分查找法
// 找到中间下标m
// 与中间数字比较，若大于中间数则 从右边数组继续二分查找
// 若数字不在数组范围内，则直接提示找不到
func find4(x int, start int, end int) {
	fmt.Println(start, end)
	if start > end || x > arrint[end] || x < arrint[start] {
		fmt.Println("找不到你要查找的数字")
		return
	}
	m := (start + end) / 2
	fmt.Printf("中间下标为%v,值为%v\n", m, arrint[m])
	if x == arrint[m] {
		fmt.Printf("找到了要查找的数字%v,下标为%v\n", x, m)
		return
	} else if x > arrint[m] {
		find4(x, m+1, end)
	} else {
		find4(x, start, m-1)
	}

}

func main() {
	fmt.Println("开始查找")
	fmt.Println("请输入要查到的信息。。。")
	// var str = ""
	// fmt.Scanln(&str)
	// find1(str)
	// find2(str)
	var x int
	fmt.Scanln(&x)
	// find3(x, arrint)
	find4(x, 0, len(arrint)-1)
}
