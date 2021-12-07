/* 搜索插入位置
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
请必须使用时间复杂度为 O(log n) 的算法。
示例 1:
输入: nums = [1,3,5,6], target = 5
输出: 2
输入: nums = [1,3,5,6], target = 7
输出: 4
*/
package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3}
	target := 2
	result := searchInsert2(nums, target)
	fmt.Println("索引为：", result)
}

// 逐个循环迭代方式
func searchInsert(nums []int, target int) int {
	flag := true
	for i := 0; i <= len(nums)-1; i++ {
		if target <= nums[i] {
			flag = false
			return i
		}
	}
	if flag {
		return len(nums)
	}
	return -1
}

func searchInsert1(nums []int, target int) int {
	for i := 0; i <= len(nums)-1; i++ {
		if target <= nums[i] {
			return i
		} else if i == len(nums)-1 {
			return i + 1
		}
	}
	return -1
}

func searchInsert2(nums []int, target int) int {

	left, right := 0, len(nums)-1
	var mid int = (left + right) / 2
	for mid >= left && mid <= right {
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
            if mid ==0 {
                return 0
            }
			if target > nums[mid-1] {
				return mid
			}
			right = mid - 1
		} else if mid <len(nums)-1 {
			if target < nums[mid +1]  {
				return mid+1
			}
			left = mid + 1
		} else {
            return len(nums)
        }
		mid = (left + right) / 2
	}
	return -1
}
