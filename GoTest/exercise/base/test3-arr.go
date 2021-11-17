package main

import "fmt"

func main() {
	// var n [10]int
	// for i := 0; i < 10; i++ {
	// 	n[i] = i + 100
	// }
	// for j := 0; j < 10; j++ {
	// 	fmt.Printf("Element[%d] = %d\n", j, n[j])
	// }
	// testarr1()
	// testarr2()
	var balance = [5]int{1000, 2, 3, 17, 50}
	var avg float32

	avg = getAverage(balance, 5)
	fmt.Printf("平均值：%f ", avg)
}

func testarr1() {
	var i, j, k int
	// 声明数组的同时快速初始化数组
	balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	for i = 0; i < 5; i++ {
		fmt.Printf("balance[%d] = %f\n", i, balance[i])
	}

	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	for j = 0; j < 5; j++ {
		fmt.Printf("balance2[%d] = %f\n", j, balance2[j])
	}

	balance3 := [5]float32{1: 2.0, 3: 7.5}
	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f\n", k, balance3[k])
	}
}

// 二维数组
func testarr2() {
	values := [][]int{}
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	values = append(values, row1)
	values = append(values, row2)

	sites := [2][2]string{}
	sites[0][0] = "Google"
	sites[0][1] = "Runoob"
	sites[1][0] = "Taobao"
	sites[1][1] = "Weibo"

	// fmt.Println("Row 1: ")
	// fmt.Println(values[0])
	// fmt.Println("Row 2: ")
	// fmt.Println(values[1])
	// fmt.Println("第3行第三个元素：")
	fmt.Println(sites)
}

func getAverage(arr [5]int, size int) float32 {
	var i, sum int
	var avg float32
	for i = 0; i < size; i++ {
		sum += arr[i]
	}
	avg = float32(sum) / float32(size)
	return avg
}
