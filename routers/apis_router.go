package routers

import (
    "github.com/astaxie/beego"
    "property-management-system/controllers"
)

func init() {
    ns := beego.NewNamespace("/api",
        beego.NSNamespace("/auths",
            beego.NSRouter("/:uid:int", &controllers.AuthController{}),
        ),
    )

    beego.AddNamespace(ns)

}

