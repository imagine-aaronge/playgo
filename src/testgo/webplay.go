package main

import (
	_ "github.com/go-sql-driver/MySQL"
    "fmt"
    "log"
    "net/http"
)

func main (){
	var db *sql.DB

}
//开启web服务
func startHttpServer() {
    http.HandleFunc("/pool", pool)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}


//db对象初始化
func init() {
	startHttpServer()
	db, _ = sql.Open("mysql","gaodun_test:414639e58d@tcp()offline.haproxy.gaodunwangxiao.com/gaodun?charset=utf8")
	db.SetMaxOpenConns(2000)
	db.setMaxIdLeConns(1000)
	db.Ping
}

func pool(w http.ReponseWriter, r *http.Request) {
	rows, err := db.Query("select * from gd_order order by id desc limit 1")
	fmt.Println(rows)
	defer rows.Close()
	checkErr(err)
}

func checkErr(err error) {
    if err != nil {
        fmt.Println(err)
        panic(err)
    }
}