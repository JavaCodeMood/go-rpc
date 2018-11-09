package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/*go连接MySQl*/
func main(){
	db,err := sql.Open("mysql","root:root@/test?charset=utf8")
	checkErr(err)

	//插入数据
	stmt,err := db.Prepare("INSERT userinfo SET username=?,department=?,created=?")
	checkErr(err)
    //传递参数
	res,err := stmt.Exec("倾城","技术部","2018-08-30")
	checkErr(err)

	id,err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	//更新数据
	stmt,err = db.Prepare("UPDATE userinfo SET username=? WHERE uid=?")
	checkErr(err)

    res,err = stmt.Exec("刘备",id)
    checkErr(err)

    affect,err := res.RowsAffected()
    checkErr(err)
    fmt.Println(affect)

    //查询数据
    rows,err := db.Query("SELECT * FROM userinfo")
    checkErr(err)

    for rows.Next(){
    	var uid int
    	var username string
    	var department string
    	var created string
    	err = rows.Scan(&uid,&username,&department,&created)
    	checkErr(err)
    	fmt.Println(uid,"-",username,"-",department,"-",created)
	}

    //删除数据
    //stmt,err = db.Prepare("DELETE FROM userinfo WHERE uid=?")
    //checkErr(err)

    //res,err = stmt.Exec(id)
    //checkErr(err)

    //affect,err = res.RowsAffected()
    //checkErr(err)

    //fmt.Println(affect)
    //关闭连接
    db.Close()
}

func checkErr(err error){
	if err != nil {
		panic(err)
	}
}
