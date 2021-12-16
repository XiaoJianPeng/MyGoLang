package main

import "fmt"

// 双指针循环
func reverseString(s []byte) {
	j := len(s) - 1
	for i := 0; i < j; i++ {
		tmp := s[i]
		s[i] = s[j]
		s[j] = tmp
		j--
	}
}

func main() {
	s := []byte("hello")
	reverseString(s)
	fmt.Println(string(s))
}
