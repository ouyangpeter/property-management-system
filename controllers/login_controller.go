package controllers

import (
    "github.com/astaxie/beego"
    "encoding/json"
    "property-management-system/responses"
)

type LoginController struct {
    beego.Controller
}

func (this *LoginController) Get() {
    userName := this.GetSession("UserName")
    if userName == nil {
        this.TplName = "login.html"
    }else {
        this.Redirect("/", 302)
    }
}

type LoginInfo struct {
    UserName    string `json:"UserName"`
    EncPassword string `json:"Password"`
}

func check(loginInfo *LoginInfo) bool {
    //TODO: check in database

    if loginInfo.UserName == "admin" {
        return true
    }else {
        return false
    }

}

func (this *LoginController) Post() {
    var loginInfo LoginInfo

    err := json.Unmarshal(this.Ctx.Input.RequestBody, &loginInfo)
    //beego.Info(string(this.Ctx.Input.RequestBody))
    if err != nil {
        // handler error 400
        res := responses.NewInvalidParameterResponse()

        res.Handler(this.Ctx.Output)
        return
    }
    if check(&loginInfo) {
        this.SetSession("UserName", loginInfo.UserName)
        this.Data["json"] = responses.NewBaseResponse()
        this.ServeJSON()

    }else {
        // handler error 401
        res := responses.NewUnauthorizedResponse()
        res.Handler(this.Ctx.Output)
        return
    }

}
