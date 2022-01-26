package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int64] int64,10)
	// 声明一个全局互斥锁， 
	lock sync.Mutex
)

func test(n int64) {
	var res int64 =1
	for i := 1; int64(i) < n; i++ {
		res *=int64(i)
	}
	 // 加锁
	lock.Lock()
	myMap[n] =res
	lock.Unlock()
}

func main() {
	for i := 1; i <= 50; i++ {
		go test(int64(i))
	}

	time.Sleep(time.Second*10)
	lock.Lock()
	for i,v :=range myMap {
		fmt.Printf("map[%d]=%d\n",i,v)
	}
	lock.Unlock()
}