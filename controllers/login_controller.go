package controllers

import (
    "github.com/astaxie/beego"
    "encoding/json"
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
    EncPassword string `json:"EncPassword"`
}

func check(loginInfo *LoginInfo) bool {
    if loginInfo.UserName == "admin" {
        return true
    }else {
        return false
    }

}

func (this *LoginController) Post() {
    var loginInfo LoginInfo

    err := json.Unmarshal(this.Ctx.Input.RequestBody, &loginInfo)
    if err != nil {
        // handler error

        return
    }
    if check(loginInfo) {
        this.SetSession("UserName", loginInfo.UserName)

    }else {
        // 403


    }

}
