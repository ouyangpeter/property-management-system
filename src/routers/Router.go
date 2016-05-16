package routers

import (
    "github.com/astaxie/beego"
    "property-management-system/src/pms"
    //"property-management-system/src/models"
)

func init() {
    //models.Initialize()

    beego.Router("/", &pms.MainController{}, "*:Index")
    publicNs := beego.NewNamespace("/public",
        beego.NSRouter("/login",
            &pms.MainController{}, "*:Login"),
        beego.NSRouter("/index",
            &pms.MainController{}, "*:Index"),
        beego.NSRouter("/logout",
            &pms.MainController{}, "*:Logout"),
        beego.NSRouter("/changepwd",
            &pms.MainController{}, "*:Changepwd"),
    )
    beego.AddNamespace(publicNs)

}