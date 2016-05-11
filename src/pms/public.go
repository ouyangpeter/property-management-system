package pms

import (
    . "property-management-system/src"
    "property-management-system/src/models"
)

type MainController struct {
    CommonController
    TemplateType string
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
    if userinfo == nil{
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
        if err == nil && id > 0{
            this.Rsp(true, "密码修改成功")
            return
        }else{
            this.Rsp(false, err.Error())
            return
        }
    }
    this.Rsp(false, "密码错误")

}
