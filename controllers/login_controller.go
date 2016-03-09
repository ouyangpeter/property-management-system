package controllers

import (
    "encoding/json"
    "github.com/astaxie/beego"
    "property-management-system/responses"
    "github.com/astaxie/beego/orm"
    "property-management-system/models"
    "strings"
)

type LoginController struct {
    beego.Controller
}

func (this *LoginController) Get() {
    userId := this.GetSession("UserId")
    if userId == nil {
        this.TplName = "login.html"
    } else {
        this.Redirect("/", 302)
    }
}

type LoginInfo struct {
    Identifier   string `json:"Identifier"`
    Credential   string `json:"Credential"`
    IdentityType string `json:"IdentityType"`
}

func (this *LoginInfo)check() (bool, uint64) {
    o := orm.NewOrm()
    userAuth := &models.UserAuths{IdentityType:this.IdentityType, Identifier:this.Identifier}
    err := o.Read(userAuth, "identity_type", "identifier")
    if err != nil {
        return false, 0
    }
    if strings.ToUpper(userAuth.Credential) == strings.ToUpper(this.Credential) {
        user := &models.User{UserId:userAuth.UserId}
        err := o.Read(user)
        if err != nil {
            return false, 0
        }

        if user.IsEnabled {
            return true, userAuth.UserId
        }
        return false, 0
    }else {
        return false, 0
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
    if ok, uid := loginInfo.check(); ok {
        this.SetSession("UserId", uid)
        this.Data["json"] = responses.NewBaseResponse()
        this.ServeJSON()

    } else {
        // handler error 401
        res := responses.NewUnauthorizedResponse()
        res.Handler(this.Ctx.Output)
        return
    }

}
