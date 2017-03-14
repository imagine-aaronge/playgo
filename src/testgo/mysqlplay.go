package main

import(
	_ "github.com/go-sql-driver/MySQL"
	"database/sql"
	"fmt"
//"time"

)


func main() {
	db,err := sql.Open("mysql","root:izumidafa123@/test?charset=utf8")//返回一个连接池对象
	 checkErr(err)
	//插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?") 
	checkErr(err)
	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	 checkErr(err)
	id, err := res.LastInsertId()
	 checkErr(err)
	fmt.Println(id)
	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)
	res, err = stmt.Exec("astaxieupdate", id)
	 checkErr(err)
	affect, err := res.RowsAffected()
	 checkErr(err)
	fmt.Println(affect)
	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo") 
	checkErr(err)
	for rows.Next() { 
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created) 
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?") 
	checkErr(err)
	res, err = stmt.Exec(id) 
	checkErr(err)
	affect, err = res.RowsAffected() 
	checkErr(err)
	fmt.Println(affect)
	db.Close()
}
func checkErr(err error) { 
	if err != nil {
	panic(err) 
	}



}
