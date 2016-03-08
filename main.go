package main

import (
    _ "property-management-system/routers"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func init() {
    beego.BConfig.CopyRequestBody = true
    beego.BConfig.WebConfig.Session.SessionOn = true
    mysqlURL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", beego.AppConfig.String("mysqluser"),
        beego.AppConfig.String("mysqlpass"), beego.AppConfig.String("mysqlurls"), beego.AppConfig.String("mysqldb"))
    orm.RegisterDataBase("default", "mysql", mysqlURL)
}

func main() {
    beego.Run()
}

