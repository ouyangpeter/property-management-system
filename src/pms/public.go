package pms

import (
    . "property-management-system/src"
    "property-management-system/src/models"
    "time"
    "github.com/astaxie/beego"
)

type MainController struct {
    CommonController
    TemplateType string
}
//community_name = xxx小区
//principal_name = 王五
//principal_phone = 13000000000
//build_date = 2014年6月8日
//parking_lot_area = 5000
//building_area = 8000
//greening_area = 1000
//road_area = 400
type CommunityInfo struct {
    CommunityName  string
    PrincipalName  string
    PrincipalPhone string
    BuildDate      string
    ParkingLotArea string
    BuildingArea  string
    GreeningArea  string
    RoadArea      string
}

var communityInfo CommunityInfo = CommunityInfo{
    CommunityName:beego.AppConfig.String("community_name"),
    PrincipalName:beego.AppConfig.String("principal_name"),
    PrincipalPhone:beego.AppConfig.String("principal_phone"),
    BuildDate:beego.AppConfig.String("build_date"),
    ParkingLotArea:beego.AppConfig.String("parking_lot_area"),
    BuildingArea:beego.AppConfig.String("building_area"),
    GreeningArea:beego.AppConfig.String("greening_area"),
    RoadArea:beego.AppConfig.String("road_area"),
}

func updateUserLoginInfo(u models.User) {
    user := models.User{Id:u.Id}
    user.LastLoginTime = time.Now()
    models.UpdateUser(&user)

}

//登录
func (this *MainController) Login() {
    isajax := this.GetString("isajax")
    if isajax == "1" {
        userName := this.GetString("username")
        passWord := this.GetString("password")
        user, err := CheckLogin(userName, passWord)
        if err == nil {
            this.SetSession("userinfo", user)
            this.Rsp(true, "登录成功")
            updateUserLoginInfo(user)
            return
        } else {
            this.Rsp(false, err.Error())
            return
        }
    }
    userinfo := this.GetSession("userinfo")
    if userinfo != nil {
        this.Ctx.Redirect(302, "/public/index")
    }
    this.TplName = this.GetTemplateType() + "/public/login.tpl"
}

func (this *MainController) Index() {
    userinfo := this.GetSession("userinfo")
    if userinfo == nil {
        this.Ctx.Redirect(302, PMS_AUTH_GATEWAY)
    }
    this.Data["userinfo"] = userinfo
    if this.GetTemplateType() != "easyui" {
        this.Layout = this.GetTemplateType() + "/public/layout.tpl"
    }
    this.TplName = this.GetTemplateType() + "/public/index.tpl"
}

func (this *MainController) Logout() {
    this.DelSession("userinfo")
    this.Ctx.Redirect(302, PMS_AUTH_GATEWAY)

}

func (this *MainController) Changepwd() {
    userinfo := this.GetSession("userinfo")
    if userinfo == nil {
        this.Ctx.Redirect(302, PMS_AUTH_GATEWAY)
    }
    oldPassword := this.GetString("oldPassword")
    newPassword := this.GetString("newPassword")
    repeatPassword := this.GetString("repeatPassword")
    if newPassword != repeatPassword {
        this.Rsp(false, "两次输入密码不一致")
    }
    user, err := CheckLogin(userinfo.(models.User).UserName, oldPassword)
    if err == nil {
        var u models.User
        u.Id = user.Id
        u.Password = newPassword
        id, err := models.UpdateUser(&u)
        if err == nil && id > 0 {
            this.Rsp(true, "密码修改成功")
            return
        } else {
            this.Rsp(false, err.Error())
            return
        }
    }
    this.Rsp(false, "密码错误")

}

func (this *MainController)Home() {
    this.Data["communityInfo"] = &communityInfo
    this.TplName = this.GetTemplateType() + "/public/home.tpl"
}
