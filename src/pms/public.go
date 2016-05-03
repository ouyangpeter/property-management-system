package pms

import (
    . "property-management-system/src"
    "github.com/astaxie/beego"
)

type MainController struct {
    CommonController
    TemplateType string
}

//登录
func (this *MainController) Login() {
    isajax := this.GetString("isajax")
    if isajax == "1"{
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
        this.Ctx.Redirect(302, beego.AppConfig.String("pms_auth_gateway"))
    }
    //if this.IsAjax()

}
