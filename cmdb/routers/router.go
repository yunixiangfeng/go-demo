package routers

import (
	"cmdb/controllers"

	"github.com/astaxie/beego"
)

// 先到内置
// 第三方
// 当前项目

func init() {
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.HomeController{})
}
