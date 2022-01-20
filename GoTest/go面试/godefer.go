package main
import "fmt"

func main() {
	for i:=0; i< 14; i++ {
		defer fmt.Println(i)
	}
}