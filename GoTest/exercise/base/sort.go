// 排序练习
package main

import (
	"fmt"
)

/* 冒泡排序
* 1. 先完成将最大的数放到最后一个位置
* 2. 将第二大的数放到倒数第二个位置
* 。。。多次循环
 */
func bubbleSort() {
	var arr = [8]int{12, 34, 65, 15, 98, 62, 37, 22}
	fmt.Println(len(arr))
	fmt.Println(arr)
	temp := 0
	for j := 0; j < len(arr)-1; j++ {
		flag := false
		for i := 0; i < len(arr)-1-j; i++ {
			if arr[i] > arr[i+1] {
				temp = arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
				flag = true
			}
		}
		fmt.Printf("第%d次排序结果arr=%v,%v\n", j+1, arr, flag)
		if !flag {
			break
		}
	}

	// for i := 0; i < len(arr); i++ {
	// 	x := arr[i]
	// 	for j := i + 1; j < len(arr); j++ {
	// 		y := arr[j]
	// 		if x < y {
	// 			x = y
	// 		}
	// 	}
	// }
	// var resultArr = []int{}

	// fmt.Println(arr)

}

func main() {
	fmt.Println("排序算法练习")
	bubbleSort()
}
