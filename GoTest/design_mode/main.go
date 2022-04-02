package main

import (
	"fmt"
	"./simplefactory"
)

func main() {
	api1 := simplefactory.NewAPI(1)
	s1 := api1.Say("Tom")
	if s1 != "Hi, Tom" {
		fmt.Println("Type1 test fail")
	} else {
		fmt.Println(s1)
	}
	api2 := simplefactory.NewAPI(2)
	s2 := api2.Say("Tom")
	if s2 != "Hello, Tom" {
		fmt.Println("Type2 test fail")
	} else {
		fmt.Println(s2)
	}
}

// //API is interface
// type API interface {
// 	Say(name string) string
// }

// //NewAPI return Api instance by type
// func NewAPI(t int) API {
// 	if t == 1 {
// 		return &hiAPI{}
// 	} else if t == 2 {
// 		return &helloAPI{}
// 	}
// 	return nil
// }

// //hiAPI is one of API implement
// type hiAPI struct{}

// //Say hi to name
// func (*hiAPI) Say(name string) string {
// 	return fmt.Sprintf("Hi, %s", name)
// }

// //HelloAPI is another API implement
// type helloAPI struct{}

// //Say hello to name
// func (*helloAPI) Say(name string) string {
// 	return fmt.Sprintf("Hello, %s", name)
// }
