/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
示例:
输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:
必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/move-zeroes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
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
