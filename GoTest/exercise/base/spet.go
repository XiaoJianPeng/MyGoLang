package main

import "fmt"

//牛顿法，开平方
func Sqrt(x float64) float64 {
	e := 1e-15
	var m, n float64
	m = x
	n = (m + x/m) / 2
	i := 0
	for m-n > e {
		m = n
		n = (m + x/m) / 2
		i++
	}
	fmt.Println("运行次数：", i)
	return m
}
func main() {
	fmt.Println(Sqrt(131))
}
