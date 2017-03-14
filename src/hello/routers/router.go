package routers

import (
    "hello/controllers"         //自身业务包
    "github.com/astaxie/beego"  //beego 包
    "github.com/beego/admin"  //admin 包
)
func init() {
     admin.Run()
    beego.Router("/", &controllers.MainController{})
}
