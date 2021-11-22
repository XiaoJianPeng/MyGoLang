package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {
	println("执行init")
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/godb")
	if err != nil {
		fmt.Println("Open mysql faild,", err)
		return
	}
	Db = database
}

func query() {
	var person []Person
	err := Db.Select(&person, "select user_id, username, sex, email from person ", )
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}
	fmt.Println("select succ:", person)
}

func insert() {
	r, err := Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	println("进入main")
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}
	fmt.Println("insert succ:", id)
}

func update() {
	res, err := Db.Exec("update person set username=? where user_id=?", "stu003", 2)
	if err != nil {
		fmt.Println("exec faild, ", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
	}
	fmt.Println("update succ:", row)
}

func delete() {
	res, err := Db.Exec("delete from person where user_id =?", 2)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
		return
	}
	fmt.Println("delete succ:", row)
}

func transaction() {
	conn, err := Db.Begin()
	if err != nil {
		fmt.Println("begin failed :", err)
		return
	}
	r, err := conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ: ", id)

	r, err = conn.Exec("insert into person(user_id, username, sex, email)values(?, ?, ?, ?)", "3","stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	id, err = r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ:", id)

	conn.Commit()
}

func main() {
	defer Db.Close()
	// 插入
	// insert()

	// 查询
	query()

	// 更新
	update()

	// 删除
	// delete()

	// 事务
	// transaction()
}
