package controllers

import (
	"cmdb/base/controllers/auth"
	"fmt"
)

type HomeController struct {
	auth.AuthorizationController
}

// Index 首页显示方法
func (c *HomeController) Index() {
	fmt.Println("Index")
	// session检查
	c.TplName = "home/index.html"
}
