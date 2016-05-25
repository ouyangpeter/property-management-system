package main

import (
    "github.com/astaxie/beego"
    "property-management-system/src/models"
    _ "property-management-system/src/routers"
    "github.com/astaxie/beego/orm"
)

func init() {
    models.Initialize()
    //beego.BConfig.CopyRequestBody = true
    //orm.RegisterDriver("mysql", orm.DRMySQL)
    //mysqlURL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", beego.AppConfig.String("mysqluser"),
    //    beego.AppConfig.String("mysqlpass"), beego.AppConfig.String("mysqlurls"), beego.AppConfig.String("mysqldb"))
    //orm.RegisterDataBase("default", "mysql", mysqlURL)
    //orm.RunSyncdb("default", false, true)

    //var filterUser = func(ctx *context.Context) {
    //    _, ok := ctx.Input.Session("UserId").(uint64)
    //    if !ok && ctx.Request.RequestURI != "/login" {
    //        ctx.Redirect(302, "/login")
    //    }
    //}
    //
    //beego.InsertFilter("/*", beego.BeforeRouter, filterUser)
}

func main() {
    orm.Debug = true
    beego.Run()
}
