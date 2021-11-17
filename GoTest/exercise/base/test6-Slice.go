/**
* 切片
 */
package main

import "fmt"

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func printSlice1() {
	numbers1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	printSlice(numbers1)
	fmt.Println("原始切片numbers ==", numbers1)
	fmt.Println("打印索引1到4 ==", numbers1[1:4])
	fmt.Println("默认下限为0 ==", numbers1[:3])
	fmt.Println("默认上限为len(s) ==", numbers1[4:])

	numbers2 := make([]int, 0, 5)
	printSlice(numbers2)
	numbers3 := numbers1[:2]
	printSlice(numbers3)

	numbers4 := numbers1[3:6]
	printSlice(numbers4)
}

func sliceCopy() {
	var numbers []int
	printSlice(numbers)
	numbers = append(numbers, 0, 1)
	printSlice(numbers)

	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	copy(numbers1, numbers)
	printSlice(numbers1)

}

func main() {
	// var numbers = make([]int, 3, 5)
	// if numbers == nil {
	// 	fmt.Println("切片是空的")
	// }
	// printSlice(numbers)
	// printSlice1()
	sliceCopy()
}
