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
            beego.NSRouter("/addBuilding", &pms.BuildingController{}, "*:Add"),
            beego.NSRouter("/deleteBuilding", &pms.BuildingController{}, "*:Delete"),
            beego.NSRouter("/updateBuilding", &pms.BuildingController{}, "*:Update"),
        ),
        beego.NSNamespace("/house",
            beego.NSRouter("/index", &pms.HouseController{}, "*:Index"),
            beego.NSRouter("/addHouse", &pms.HouseController{}, "*:Add"),
            beego.NSRouter("/deleteHouse", &pms.HouseController{}, "*:Delete"),
            beego.NSRouter("/updateHouse", &pms.HouseController{}, "*:Update"),
        ),
    )
    beego.AddNamespace(publicNs)
    beego.AddNamespace(pmsNs)

}