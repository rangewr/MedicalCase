package main

import (
	_ "MedicalCase/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//************************临时测试接口使用,解决跨域问题***************************
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH,POST, GET, OPTIONS, DELETE"},
		AllowHeaders:     []string{"Origin,Content-Type,XFILENAME,XFILECATEGORY,XFILESIZE,x-requested-with"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true}))
	//********************************************************************************
	beego.Run()
}

//初始化orm,并设置数据库连接
func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	connectionString := "root:123456@tcp(localhost:3306)/kstore?charset=utf8"
	orm.RegisterDataBase("default", "mysql", connectionString)
	orm.Debug = true
}
