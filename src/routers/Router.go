package routers

import (
    "github.com/astaxie/beego"
    "property-management-system/src/pms"
    //"property-management-system/src/models"
)

func init() {
    //models.Initialize()

    publicNs := beego.NewNamespace("/public",
        beego.NSRouter("/login",
            &pms.MainController{}, "*:Login"),
        beego.NSRouter("/index",
            &pms.MainController{}, "*:Index"),
    )
    beego.AddNamespace(publicNs)

}