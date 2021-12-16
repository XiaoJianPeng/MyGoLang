/*
给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
示例 1:
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

/*
 使用slice特性，先把需要数组需要移动的后几位截取，再与前部组成新的切片，copy给nums
*/
func rotate(nums []int, k int) {
	l := len(nums)
	if l <= k {
		k = k % l
	}
	if k > 0 {
		movearr := nums[l-k : l]
		movearr = append(movearr, nums[:l-k]...)
		copy(nums, movearr)
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)

}
