package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	workerBits uint8 = 10
	numberBits uint8 = 12
	// ^异或运算； << 二进制左移n位，十进制就是乘以2的n次方，
	workerMax   int64 = -1 ^ (-1 << workerBits)
	numberMax   int64 = -1 ^ (-1 << numberBits)
	timeShift   uint8 = workerBits + numberBits
	workerShift uint8 = numberBits
	startTime   int64 = 1641884394126 // 如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID
)

type Worker struct {
	// Go语言互斥锁 Mutex ；读写互斥锁（sync.RWMutex）
	mu        sync.Mutex
	timestamp int64
	workerId  int64
	number    int64
}

func NewWorker(workerId int64) (*Worker, error) {
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("Worker ID excess of quantity")
	}
	return &Worker{
		timestamp: 0,
		workerId:  workerId,
		number:    0,
	}, nil
}

// 生成id
func (w *Worker) GetId() int64 {
	// 加锁
	w.mu.Lock()
	// 延迟释放锁
	defer w.mu.Unlock()
	now := time.Now().UnixNano() / 1e6
	if w.timestamp == now {
		w.number++
		if w.number > numberMax {
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.number = 0
		w.timestamp = now
	}
	// | 二进位相或
	ID := int64((now-startTime)<<int64(timeShift) | (w.workerId << int64(workerShift)) | w.number)
	return ID
}

func main() {
	// fmt.Println(-1 << 10)
	// fmt.Println(1 ^ 1024)
	// fmt.Println(-1 << 12)
	// fmt.Println(-1 ^ -4096)
	// fmt.Println("workerBits:====>", workerBits)
	// fmt.Println("numberBits:====>", numberBits)
	// fmt.Println("workerMax:====>", workerMax)
	// fmt.Println("numberMax:====>", numberMax)
	// fmt.Println("timeShift:====>", timeShift)
	// fmt.Println("workerShift:====>", workerShift)
	startTime := time.Now().Local().UnixMilli()
	fmt.Println(startTime)
	node, err := NewWorker(1)
	if err != nil {
		panic(err)
	}
	for {
		fmt.Println(node.GetId())
	}
}
