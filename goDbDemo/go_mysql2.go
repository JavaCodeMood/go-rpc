package main

import (
   "database/sql"
   "fmt"
   _ "github.com/go-sql-driver/mysql"
   "log"
   "net/http"
)

var db *sql.DB

//DB对象初始化
func init() {
	//设置驱动和连接
    db, _ = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_db?charset=utf8")
    db.SetMaxOpenConns(2000)  //用于设置最大打开的连接数，默认值为0表示不限制。
    db.SetMaxIdleConns(1000)  //用于设置闲置的连接数。
    db.Ping()
}

func main() {
   startHttpServer()
}

//启动一个web服务监听9090端口
func startHttpServer() {
	http.HandleFunc("/pool", pool)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
       log.Fatal("ListenAndServe: ", err)
    }
}

func pool(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT * FROM users ")
    defer rows.Close()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(map[string]string)
	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
			fmt.Println(i)
		}
    }

	fmt.Println(record)
	fmt.Fprintln(w, "finish")
}

