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
    pmsNs := beego.NewNamespace("/pms",
        beego.NSNamespace("/building",
            beego.NSRouter("/index", &pms.BuildingController{}, "*:Index"),
            beego.NSRouter("/addBuilding", &pms.BuildingController{}, "*:AddBuilding"),
            beego.NSRouter("/deleteBuilding", &pms.BuildingController{}, "*:DeleteBuilding"),
            beego.NSRouter("/updateBuilding", &pms.BuildingController{}, "*:UpdateBuilding"),
        ),
    )
    beego.AddNamespace(publicNs)
    beego.AddNamespace(pmsNs)

}