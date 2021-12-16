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
