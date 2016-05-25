package pms

import (
    m "property-management-system/src/models"
)

type OwnerController struct {
    CommonController
}

func (this *OwnerController)Index() {
    page, _ := this.GetInt64("page")
    page_size, _ := this.GetInt64("rows")
    sort := this.GetString("sort")
    order := this.GetString("order")

    queryData := m.OwnerQueryParam{
    }

    if len(order) > 0 {
        if order == "desc" {
            sort = "-" + sort
        }
    } else {
        sort = "Id"
    }
    owners, count := m.GetOwnerList(page, page_size, sort, queryData)
    for _, owner := range owners {
        owner.User.Owner = nil;
    }

    if this.IsAjax() {
        if owners == nil {
            owners = make([]m.Owner, 0)
        }
        this.Data["json"] = &map[string]interface{}{"total":count, "rows":owners}
        //log.Println(houses[0].Building)
        this.ServeJSON()
        return
    } else {
        buildings, _ := m.GetAllBuilding()
        this.Data["buildings"] = &buildings
        this.Data["owners"] = &owners
        if this.GetTemplateType() != "easyui" {
            this.Layout = this.GetTemplateType() + "/public/layout.tpl"
        }
        this.TplName = this.GetTemplateType() + "/pms/owner.tpl"
    }

}

func (this *OwnerController)Add() {
    owner := m.Owner{}
    if err := this.ParseForm(&owner); err != nil {
        this.Rsp(false, err.Error())
        return
    }
    id, err := m.AddOwner(&owner)
    if id > 0 && err == nil {
        this.Rsp(true, "Success")
        return
    } else {
        this.Rsp(false, err.Error())
        return
    }

}

func (this *OwnerController)Delete() {
    Id, _ := this.GetInt64("Id")
    status, err := m.DeleteOwnerById(Id)
    if status > 0 && err == nil {
        this.Rsp(true, "Success")
        return
    } else {
        this.Rsp(false, err.Error())
        return
    }
}