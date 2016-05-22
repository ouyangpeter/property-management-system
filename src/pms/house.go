package pms

import (
    m "property-management-system/src/models"
)

type HouseController struct {
    CommonController
}

func (this *HouseController)Index() {
    page, _ := this.GetInt64("page")
    page_size, _ := this.GetInt64("rows")
    sort := this.GetString("sort")
    order := this.GetString("order")
    area, _ := this.GetInt("area")
    houseNo := this.GetString("house_no")
    unitName := this.GetString("unit_name")
    buildingId, _ := this.GetInt64("building_id")
    ownerId, _ := this.GetInt64("owner_id")

    queryData := m.HouseQueryParam{
        HouseNo:houseNo,
        Area: area,
        UnitName:unitName,
        OwnerId:ownerId,
        BuildingId:buildingId,
    }

    if len(order) > 0 {
        if order == "desc" {
            sort = "-" + sort
        }
    } else {
        sort = "Id"
    }
    houses, count := m.GetHouseList(page, page_size, sort, queryData)

    if this.IsAjax() {
        if houses == nil {
            houses = make([]m.House, 0)
        }
        this.Data["json"] = &map[string]interface{}{"total":count, "rows":houses}
        //log.Println(houses[0].Building)
        this.ServeJSON()
        return
    } else {
        buildings, _ := m.GetAllBuilding()
        this.Data["buildings"] = &buildings
        this.Data["houses"] = &houses
        if this.GetTemplateType() != "easyui" {
            this.Layout = this.GetTemplateType() + "/public/layout.tpl"
        }
        this.TplName = this.GetTemplateType() + "/pms/house.tpl"
    }

}

func (this *HouseController)Add() {
    house := m.House{}
    if err := this.ParseForm(&house); err != nil {
        this.Rsp(false, err.Error())
        return
    }
    id, err := m.AddHouse(&house)
    if id > 0 && err == nil {
        this.Rsp(true, "Success")
        return
    } else {
        this.Rsp(false, err.Error())
        return
    }
}

func (this *HouseController)Delete() {
    Id, _ := this.GetInt64("Id")
    status, err := m.DeleteHouseById(Id)
    if status > 0 && err == nil {
        this.Rsp(true, "Success")
        return
    } else {
        this.Rsp(false, err.Error())
        return
    }
}

func (this *HouseController)Update() {
    house := m.House{}
    if err := this.ParseForm(&house); err != nil {
        this.Rsp(false, err.Error())
        return
    }
    id, err := m.UpdateHouse(&house)
    if err == nil && id > 0 {
        this.Rsp(true, "Success")
        return
    } else {
        this.Rsp(false, err.Error())
        return
    }

}
