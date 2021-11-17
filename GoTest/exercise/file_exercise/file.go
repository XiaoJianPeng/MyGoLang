package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 打开文件
func openFile() {
	file, err := os.Open("d:/test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	} else {
		fmt.Printf("file=%v", file)
	}
	err = file.Close()
	if err != nil {
		fmt.Println("close file err", err)
	}
}

// 使用ioutil.ReadFile 读取文件,无需close()  content为[]byte
func readFile1(file string) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file eerr =%v", err)
	}
	// 展示读取到的内容
	fmt.Printf("%v", string(content))
}

// 读取文件，适合大文件
func rederFile2() {
	file, err := os.Open("d:/test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	} else {
		fmt.Printf("file=%v\n", file)
	}
	//带缓冲区
	defer file.Close()
	// file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束。")
}

// 创建一个新文件,并写入内容
func createFile() {
	filepath := "d:/abc.txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()
	// 准备写入5句话
	str := "hello, Gardon\n"
	// 写入时， 使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	// 因为Writer是写入缓存的需要先保存到文件里
	writer.Flush()
	readFile1(filepath)
}

// 打开一个文件覆盖文件内容O_TRUNC
func rewriteFile() {
	filepath := "d:/abc.txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()
	str := "你好，地球\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
	readFile1(filepath)
}

// 打开一个已存在的文件，追加文件内容，
func modifyFile() {
	filepath := "d:/abc.txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()
	str := "abcdefge,航哈户撒，打碎怒闯撒\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
	readFile1(filepath)
}

// 打开已存在文件，将文件内容输出到中断，并追加内容
func readWriteFile() {
	filepath := "d:/abc.txt"
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束。")
	str := "读写\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}

// 判断文件或目录是否存在
func fileExists() {
	filepath := "d:/test.txt"
	file, err := os.Stat(filepath)
	if err == nil {
		fmt.Println("文件存在:", file)
	} else {
		fmt.Printf("文件不存在，err=%v", err.Error())
	}
}

// 文件拷贝
func copyFile() {

	filepath1 := "C:/Users/DELL/Pictures/微信图片_20210318190511.jpg"
	filepath2 := "D:/迁移图片.jpg"

	srcFile, err := os.Open(filepath1)
	if err != nil {
		fmt.Println("文件不存在")
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	distFile, err := os.OpenFile(filepath2, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("err=", err.Error())
		return
	}
	defer distFile.Close()
	writer := bufio.NewWriter(distFile)
	io.Copy(writer, reader)
}

func main() {
	copyFile()
}
