package main

import "fmt"

// slice 应用
func moveZeroes1(nums []int) {
	slice0 := make([]int, 0)
	slice1 := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		fmt.Println(i)
		if nums[i] == 0 {
			slice0 = append(slice0, nums[i])
		} else {
			slice1 = append(slice1, nums[i])
		}
	}
	slice1 = append(slice1, slice0...)
	copy(nums, slice1)
}

// 数组循环
func moveZeroes2(nums []int) {
	j := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for i := 0; i < n-j; i++ {
		nums[n-i-1] = 0
	}

}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes2(nums)
	fmt.Println("结果输出：", nums)
}
