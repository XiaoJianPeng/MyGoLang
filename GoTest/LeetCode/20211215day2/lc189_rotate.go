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
