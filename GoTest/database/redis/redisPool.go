// redis 连接池
package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	// 实例化一个连接池
	pool = &redis.Pool{
		// 最初的连接数量
		MaxIdle: 16,
		// 最大连接数， 0代表自动定义，按需分配
		MaxActive: 0,
		// 连接关闭时间300秒，300秒不适用自动关闭
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	//从连接池娶一个链接
	c := pool.Get()
	defer c.Close()

	_, err := c.Do("Set", "qwe", 210)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Int(c.Do("Get", "qwe"))
	if err != nil {
		fmt.Println("get qwe failed: ", err)
		return
	}
	fmt.Println(r)
	pool.Close()
}
