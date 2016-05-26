package pms

import (
    m "property-management-system/src/models"
)

type OwnerController struct {
    CommonController
}

func (this *OwnerController)Index() {
    //building_id:101
    //unit_name:2
    //house_id:1
    //owner_name:123
    //owner_phone:456
    //owner_idcard:789
    //owner_company:10
    //page:1
    //rows:20
    page, _ := this.GetInt64("page")
    page_size, _ := this.GetInt64("rows")
    sort := this.GetString("sort")
    order := this.GetString("order")
    buildingId, _ := this.GetInt64("building_id")
    unitName := this.GetString("unit_name")
    houseId, _ := this.GetInt64("house_id")
    ownerName := this.GetString("owner_name")
    ownerPhone := this.GetString("owner_phone")
    ownerIdCard := this.GetString("owner_idcard")
    ownerCompany := this.GetString("owner_company")

    queryData := m.OwnerQueryParam{
        Name:ownerName,
        BuildingId:buildingId,
        UnitName:unitName,
        HouseId:houseId,
        PhoneNumber:ownerPhone,
        IdCard:ownerIdCard,
        Company:ownerCompany,
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

func (this *OwnerController)Update() {
    owner := m.Owner{}
    if err := this.ParseForm(&owner); err != nil {
        this.Rsp(false, err.Error())
        return
    }
    id, err := m.UpdateOwner(&owner)
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