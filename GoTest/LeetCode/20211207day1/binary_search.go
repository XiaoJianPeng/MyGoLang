/** 二分查找
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
示例 1:
输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4
示例 2:
输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	var result int
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 2
	fmt.Println(time.Now())
	// result = search1(nums, target)
	result = search2(nums, target) // 137,474 ;139,750
	fmt.Println(time.Now())
	fmt.Println("输出结果：", result)
}

func search1(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}
func search2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		if target > nums[right] || target < nums[left] {
			return -1
		}
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			left = mid + 1
			continue
		} else {
			right = mid - 1
			continue
		}
	}
	return -1
}
