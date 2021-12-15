package main

import "fmt"

func sortedSquares(nums []int) []int {
	n :=len(nums)-1
	ans := make([]int, n+1)
	i, j, k := 0, n, n
	for i<=j {
		lm, rm := nums[i]*nums[i], nums[j]*nums[j]
		if lm > rm {
			ans[k] =lm
			i++
		} else {
			ans[k]=rm
			j--
		}
		k--
	}
	return ans
}

func main() {
	arr := []int{-12, -7, -3, 0, 4, 5, 8, 10}
	res := sortedSquares(arr)
	fmt.Println("结果res：", res)
}
