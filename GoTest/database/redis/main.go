package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var conn redis.Conn

func init() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	conn = c
	fmt.Println("redis conn success")
}

func main() {
	defer conn.Close()
	redis_set("ab", "208")
	redis_get("ab")
	mset()
	mget()
	redis_expire()
	redis_list()
	redis_hash()
}

// 设置过期时间
func redis_expire() {
	_, err := conn.Do("expire", "abc", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// set key
func redis_set(key string, value string) {
	_, err := conn.Do("Set", key, value)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// get
func redis_get(key string) {
	r, err := redis.Int(conn.Do("Get", key))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}
	fmt.Println(r)
}

// 批量set
func mset() {
	_, err := conn.Do("MSet", "wwwf", 100, "efg", 300)
	if err != nil {
		fmt.Println(err)
		return
	}
}

//批量get
func mget() {
	r, err := redis.Strings(conn.Do("Mget", "abc", "wwwf", "efg"))
	if err != nil {
		fmt.Println("Mget failed,", err)
		return
	}
	for _, v := range r {
		fmt.Println(v)
	}
}

// List队列操作, 栈内数据取完后，栈消失
func redis_list() {
	_, err := conn.Do("lpush", "book_list", "abc", "ceg", 300)
	if err != nil {
		fmt.Println("err===>", err)
		return
	}

	r, err := redis.String(conn.Do("lpop", "book_list"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}
	fmt.Println(r)
}

// hash表
func redis_hash() {
	_, err := conn.Do("HSet", "books", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	r, err := redis.Int(conn.Do("HGet", "books", "abc"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}
	fmt.Println(r)
}
