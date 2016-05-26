package pms

import (
    m "property-management-system/src/models"
    "github.com/astaxie/beego/orm"
)

type UserController struct {
    CommonController
}

func (this *UserController)Index() {
    page, _ := this.GetInt64("page")
    page_size, _ := this.GetInt64("rows")
    sort := this.GetString("sort")
    order := this.GetString("order")
    userName := this.GetString("user_name")
    email := this.GetString("email")
    status, _ := this.GetInt("status")

    queryData := m.UserQueryParam{
        Type:12,
        UserName:userName,
        Email:email,
        Status:status,
    }

    if len(order) > 0 {
        if order == "desc" {
            sort = "-" + sort
        }
    } else {
        sort = "Id"
    }
    users, count := m.GetUserList(page, page_size, sort, queryData)

    if this.IsAjax() {
        if users == nil {
            users = make([]orm.Params, 0)
        }
        this.Data["json"] = &map[string]interface{}{"total":count, "rows":users}
        //log.Println(houses[0].Building)
        this.ServeJSON()
        return
    } else {
        this.Data["users"] = &users
        if this.GetTemplateType() != "easyui" {
            this.Layout = this.GetTemplateType() + "/public/layout.tpl"
        }
        this.TplName = this.GetTemplateType() + "/pms/user.tpl"
    }

}

func (this *UserController)Add() {
    user := m.User{}
    if err := this.ParseForm(&user); err != nil {
        this.Rsp(false, err.Error())
        return
    }
    //add管理员
    user.Type = 12
    id, err := m.AddUser(&user)
    if id > 0 && err == nil {
        this.Rsp(true, "Success")
        return
    } else {
        this.Rsp(false, err.Error())
        return
    }

}

func (this *UserController)Delete() {
    Id, _ := this.GetInt64("Id")
    status, err := m.DeleteUserById(Id)
    if status > 0 && err == nil {
        this.Rsp(true, "Success")
        return
    } else {
        this.Rsp(false, err.Error())
        return
    }

}

func (this *UserController)Update() {
    user := m.User{}
    if err := this.ParseForm(&user); err != nil {
        this.Rsp(false, err.Error())
        return
    }
    id, err := m.UpdateUser(&user)
    if err == nil && id > 0 {
        this.Rsp(true, "Success")
        return
    } else if err == nil {
        this.Rsp(false, "Nothing has been updated")
        return
    } else {
        this.Rsp(false, err.Error())
        return
    }
}
