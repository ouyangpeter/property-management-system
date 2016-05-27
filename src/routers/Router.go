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
            beego.NSRouter("/buildingList", &pms.BuildingController{}, "*:GetAllBuildingList"),
            beego.NSRouter("/unitList", &pms.BuildingController{}, "*:GetUnitListByBuildingId"),
        ),
        beego.NSNamespace("/house",
            beego.NSRouter("/index", &pms.HouseController{}, "*:Index"),
            beego.NSRouter("/addHouse", &pms.HouseController{}, "*:Add"),
            beego.NSRouter("/deleteHouse", &pms.HouseController{}, "*:Delete"),
            beego.NSRouter("/updateHouse", &pms.HouseController{}, "*:Update"),
            beego.NSRouter("/houseList", &pms.HouseController{}, "*:GetHouseList"),
        ),
        beego.NSNamespace("/owner",
            beego.NSRouter("/index", &pms.OwnerController{}, "*:Index"),
            beego.NSRouter("/addOwner", &pms.OwnerController{}, "*:Add"),
            beego.NSRouter("/deleteOwner", &pms.OwnerController{}, "*:Delete"),
            beego.NSRouter("/updateOwner", &pms.OwnerController{}, "*:Update"),
        ),
        beego.NSNamespace("/user",
            beego.NSRouter("/index", &pms.UserController{}, "*:Index"),
            beego.NSRouter("/addUser", &pms.UserController{}, "*:Add"),
            beego.NSRouter("/deleteUser", &pms.UserController{}, "*:Delete"),
            beego.NSRouter("/updateUser", &pms.UserController{}, "*:Update"),
        ),
        beego.NSNamespace("/parkingLot",
            beego.NSRouter("/index", &pms.ParkingLotController{}, "*:Index"),
            beego.NSRouter("/addParkingLot", &pms.ParkingLotController{}, "*:Add"),
            beego.NSRouter("/deleteParkingLot", &pms.ParkingLotController{}, "*:Delete"),
            beego.NSRouter("/updateParkingLot", &pms.ParkingLotController{}, "*:Update"),
            beego.NSRouter("/parkingLotList", &pms.ParkingLotController{}, "*:GetAllParkingLotList"),
        ),
        beego.NSNamespace("/parkingSpot",
            beego.NSRouter("/index", &pms.ParkingSpotController{}, "*:Index"),
            beego.NSRouter("/addParkingSpot", &pms.ParkingSpotController{}, "*:Add"),
            beego.NSRouter("/deleteParkingSpot", &pms.ParkingSpotController{}, "*:Delete"),
            beego.NSRouter("/updateParkingSpot", &pms.ParkingSpotController{}, "*:Update"),
        ),
    )
    beego.AddNamespace(publicNs)
    beego.AddNamespace(pmsNs)

}