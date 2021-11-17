// 结构体
package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func initBook(title, author, subject string, book_id int) Books {
	var Book Books
	Book.title = title
	Book.author = author
	Book.subject = subject
	Book.book_id = book_id
	return Book
}

func printBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func main() {
	// fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

	// fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 67231})

	var Book1, Book2 Books
	Book1 = initBook("Go 语言", "www.runoob.com", "Go 语言教程", 6495407)
	Book2 = initBook("Python 教程", "www.runoob.com", "Python 语言教程", 6495700)
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.book_id)

	/* 打印 Book2 信息 */
	fmt.Print("Book 2 :", Book2)

	printBook(&Book1)
}
