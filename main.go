package main

import (
    "fmt"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
    _ "property-management-system/routers"
)

func init() {
    beego.BConfig.CopyRequestBody = true
    beego.BConfig.WebConfig.Session.SessionOn = true
    orm.RegisterDriver("mysql", orm.DRMySQL)
    mysqlURL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", beego.AppConfig.String("mysqluser"),
        beego.AppConfig.String("mysqlpass"), beego.AppConfig.String("mysqlurls"), beego.AppConfig.String("mysqldb"))
    orm.RegisterDataBase("default", "mysql", mysqlURL)
    orm.RunSyncdb("default", false, true)
}

func main() {
    beego.Run()
}
