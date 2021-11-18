// 命令行参数
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("命令行参数有：%v", len(os.Args))
	// for i, v := range os.Args {
	// 	fmt.Printf("args[%v]=== %v\n", i, v)
	// }
}
