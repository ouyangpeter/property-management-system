package controllers

import (
    "github.com/astaxie/beego"
    "property-management-system/models"
    "encoding/json"
    "property-management-system/responses"
    "github.com/astaxie/beego/orm"
    "strconv"
)

type AuthController struct {
    beego.Controller
}

func isUserPassword(identityType string) (bool) {
    if identityType == "username" || identityType == "phone" || identityType == "email" {
        return true
    }
    return false
}

func (this *AuthController) Patch() {
    var userAuth models.UserAuth

    err := json.Unmarshal(this.Ctx.Input.RequestBody, &userAuth)

    if err != nil {
        res := responses.NewInvalidParameterResponse()
        res.Handler(this.Ctx.Output)
        return
    }
    // TODO if uid != this.GetSession("UserId")
    tmpUid := this.Ctx.Input.Param(":uid")
    tmpSessionUid := this.GetSession("UserId")

    sessionUid, ok := tmpSessionUid.(uint64)
    beego.Info("sessionUid", sessionUid)
    if !ok {
        res := responses.NewInternalErrorResponse()
        res.Handler(this.Ctx.Output)
        return
    }

    uid, err := strconv.ParseUint(tmpUid, 10, 64)

    if err != nil {
        res := responses.NewInternalErrorResponse()
        res.Handler(this.Ctx.Output)
        return
    }

    beego.Info("tmpUid:%s, sessionUid:%s", tmpUid, sessionUid)

    // 要patch的uid和登录者不是同一个
    if uid != sessionUid {
        res := responses.NewForbiddenResponse()
        res.Handler(this.Ctx.Output)
        return
    }

    // IdentityType是username phone email 的都要修改
    beego.Info(userAuth)
    if isUserPassword(userAuth.IdentityType) {
        userAuth.User.UserId = uid
        o := orm.NewOrm()

        _, err := o.QueryTable("UserAuth").Filter("UserId", userAuth.User.UserId).Filter("IdentityType__in", "username", "phone", "email").Update(orm.Params{"Credential":userAuth.Credential})

        if err != nil {
            res := responses.NewInternalErrorResponse()
            res.Handler(this.Ctx.Output)
            return
        }

        this.Data["json"] = responses.NewBaseResponse()
        this.ServeJSON()
    }else {
        res := responses.NewInvalidParameterResponse()
        res.Handler(this.Ctx.Output)
        return
    }

}

