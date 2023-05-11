package controllers

import (
	"cmdb/base/errors"
	"cmdb/forms"
	"cmdb/models"

	"net/http"

	"github.com/astaxie/beego"
)

// AuthController 认证控制器
type AuthController struct {
	beego.Controller
}

// Login 认证登录
func (c *AuthController) Login() {
	form := &forms.LoginForm{}
	errs := errors.New()
	// Get请求直接加载页面
	// Post请求进行数据验证
	if c.Ctx.Input.IsPost() {
		// 获取用户提交数据
		if err := c.ParseForm(form); err == nil {
			user := models.GetUserByName(form.Name)

			if user == nil {
				errs.Add("default", "用户名或密码错误")
				// 用户不存在
			} else if user.ValidPassword(form.Password) {
				c.Redirect("/home/index", http.StatusFound)
			} else {
				// 用户密码不正确
				errs.Add("default", "用户名或密码错误")
			}
		} else {
			errs.Add("default", "用户名或密码错误")
		}
	}

	c.Data["form"] = form
	c.Data["errors"] = errs
	// 定义加载页面
	c.TplName = "auth/login.html"

}
