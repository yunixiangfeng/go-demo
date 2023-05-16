package main

import (
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"

	_ "cmdb/routers"
)

func main() {
	orm.Debug = true
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/cmdb?charset=utf8mb4&loc=PRC&parseTime=true",
		beego.AppConfig.DefaultString("mysql::User", "root"),
		beego.AppConfig.DefaultString("mysql::Password", "1234"),
		beego.AppConfig.DefaultString("mysql::Host", "192.168.204.130"),
		beego.AppConfig.DefaultInt("mysql::Port", 3306),
	)

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dsn)

	// 测试数据库连接
	if db, err := orm.GetDB("default"); err != nil {
		log.Fatal(err)
	} else if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	beego.Run()
}
