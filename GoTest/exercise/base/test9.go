// Go递归
package main

import "fmt"

func factorial(n uint64) (result uint64) {
	if n > 0 {
		fmt.Printf("当前n%v\n", n)
		result = n * factorial(n-1)
		return result
	}
	return 1
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	i := 10
	fmt.Println(factorial(uint64(i)))
	for i := 0; i < 10; i++ {
		fmt.Printf("%d,结果%d\n", i, fibonacci(i))
	}
}
