package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int
	Sal   float64
	Skill string
}

// 将struct序列化
func testStruct() {
	monster := Monster{
		Name:  "洪七公",
		Age:   55,
		Sal:   12000.00,
		Skill: "降龙十八掌",
	}

	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化错误 err=&v", err)
	}
	fmt.Printf("序列化后的结果%v", string(data))
}

// 将一个Map序列化
func testMap() {
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["name"] = "张三丰"
	m["age"] = 100
	m["address"] = "武当山"

	data, err := json.Marshal(&m)
	if err != nil {
		fmt.Println("序列化错误 err=&v", err)
	}
	fmt.Printf("序列化后的结果%v", string(data))
}

func testArr() {
	arr := [...]Monster{
		{
			Name:  "洪七公",
			Age:   55,
			Sal:   12000.00,
			Skill: "降龙十八掌",
		},
		{
			Name:  "王重阳",
			Age:   65,
			Sal:   16000.00,
			Skill: "重阳先天功",
		},
		{
			Name:  "黄药师",
			Age:   52,
			Sal:   16000.00,
			Skill: "弹指神通",
		},
	}

	data, err := json.Marshal(&arr)
	if err != nil {
		fmt.Println("序列化错误 err=&v", err)
	}
	fmt.Printf("序列化后的结果%v", string(data))
}

// 反序列化
func unmarshalArr() {
	str := "{\"Name\":\"洪七公\",\"Age\":55,\"Sal\":12000,\"Skill\":\"降龙十八掌\"}"
	var m Monster
	err := json.Unmarshal([]byte(str), &m)
	if err == nil {
		fmt.Println(m)
	}
}

// 反序列化成切片
func unmarshalSlice() {
	str := "[{\"Name\":\"洪七公\",\"Age\":55,\"Sal\":12000,\"Skill\":\"降龙十八掌\"}," +
		"{\"Name\":\"王重阳\",\"Age\":65,\"Sal\":16000,\"Skill\":\"重阳先天功\"}]"

	var slice []map[string]interface{}
	// 反序列化，map不需要make操作，因为make操作被封装到 Unmarshal函数中了
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("err====%v\n", err)
	}
	fmt.Printf("slice反序列化后结果slice=%v\n", slice)
}

func main() {
	unmarshalSlice()
}
