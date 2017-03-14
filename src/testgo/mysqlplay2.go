package main

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
//"time"

)
var db,err = sql.Open("mysql","gaodun_test:414639e58d@tcp(offline.haproxy.gaodunwangxiao.com:3307)/gaodun?charset=utf8")

func main(){
	checkErr(err)
	// query()
	update()
}


func query() {
	rows, err := db.Query("select paytime from gd_order where id = 1234")
	checkErr(err)

	for rows.Next() {
		var paytime int
		rows.Columns()
		err = rows.Scan(&paytime)
		checkErr(err)
		fmt.Println(paytime)
	}
}

func update() {
	stmt, err := db.Prepare(`update gd_order set paytime =? where id = ?`)
	checkErr(err)
	res, err := stmt.Exec(13917607695,1234)
	checkErr(err)
	num, err :=res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}


func insert() {
	
}

func checkErr(err error) { 
	if err != nil {
	panic(err) 
	}



}
