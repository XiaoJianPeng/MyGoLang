/*
给定一个已按照 非递减顺序排列  的整数数组 numbers ，请你从数组中找出两个数满足相加之和等于目标数 target 。

函数应该以长度为 2 的整数数组的形式返回这两个数的下标值。numbers 的下标 从 1 开始计数 ，所以答案数组应当满足 1 <= answer[0] < answer[1] <= numbers.length 。

你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。
示例 1：
输入：numbers = [2,7,11,15], target = 9
输出：[1,2]
解释：2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。
示例 2：
输入：numbers = [2,3,4], target = 6
输出：[1,3]
示例 3：
输入：numbers = [-1,0], target = -1
输出：[1,2]
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

// 此方法在极限数组情况下运行超时
func twoSum1(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > target {
			return nil
		} else {
			othernum := target - numbers[i]
			for j := i + 1; j < len(numbers); j++ {
				fmt.Println("j=", j)
				if numbers[j] > othernum {
					break
				} else if numbers[j] == othernum {
					return []int{i + 1, j + 1}
				}
			}
		}
	}
	return nil
}

// 双指针模式，从两端开始相加比较，向中间靠拢
func twoSum2(numbers []int, target int) []int {
	j := len(numbers) - 1
	for i := 0; i < j; {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum < target {
			i++
		} else {
			j--
		}
	}
	return nil
}

func main() {
	numbers := []int{0, 2, 7, 11, 15}
	result := twoSum2(numbers, 9)
	fmt.Println("result===", result)
}
