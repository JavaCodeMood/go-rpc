package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

)


func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_db?charset=utf8")
	checkErrs(err)

	//插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErrs(err)

	res, err := stmt.Exec("码农", "研发部门", "2016-03-06")
	checkErrs(err)

	id, err := res.LastInsertId()
	checkErrs(err)

	fmt.Println(id)
	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErrs(err)

	res, err = stmt.Exec("码农二代", id)
	checkErrs(err)

	affect, err := res.RowsAffected()
	checkErrs(err)

	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErrs(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErrs(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErrs(err)

	res, err = stmt.Exec(id)
	checkErrs(err)

	affect, err = res.RowsAffected()
	checkErrs(err)

	fmt.Println(affect)

	db.Close()

}

func checkErrs(err error) {
	if err != nil {
		fmt.Println("发生错误：", err)
	}
}

