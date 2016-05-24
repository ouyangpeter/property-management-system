package pms

import (
    m "property-management-system/src/models"
    "github.com/astaxie/beego/orm"
)

type BuildingController struct {
    CommonController
}

func (this *BuildingController)Index() {
    page, _ := this.GetInt64("page")
    page_size, _ := this.GetInt64("rows")
    sort := this.GetString("sort")
    order := this.GetString("order")
    name := this.GetString("name")
    floors, _ := this.GetInt("floors")
    height, _ := this.GetInt("height")
    area, _ := this.GetInt("area")
    queryData := m.BuildingQueryParam{
        Name:name,
        Floors:floors,
        Area: area,
        Height:height,
    }

    if len(order) > 0 {
        if order == "desc" {
            sort = "-" + sort
        }
    } else {
        sort = "Id"
    }
    buildings, count := m.GetBuildingList(page, page_size, sort, queryData)

    if this.IsAjax() {
        if buildings == nil {
            buildings = make([]orm.Params, 0)
        }
        this.Data["json"] = &map[string]interface{}{"total":count, "rows":buildings}
        this.ServeJSON()
        return
    } else {
        this.Data["building"] = &buildings
        if this.GetTemplateType() != "easyui" {
            this.Layout = this.GetTemplateType() + "/public/layout.tpl"
        }
        this.TplName = this.GetTemplateType() + "/pms/building.tpl"
    }
}

func (this *BuildingController)Add() {
    storiedBuilding := m.StoriedBuilding{}
    if err := this.ParseForm(&storiedBuilding); err != nil {
        this.Rsp(false, err.Error())
        return
    }
    id, err := m.AddBuilding(&storiedBuilding)
    if id > 0 && err == nil {
        this.Rsp(true, "Success")
        return
    } else {
        this.Rsp(false, err.Error())
        return
    }
}

func (this *BuildingController) Delete() {
    Id, _ := this.GetInt64("Id")
    status, err := m.DeleteBuildingById(Id)
    if status > 0 && err == nil {
        this.Rsp(true, "Success")
        return
    } else {
        this.Rsp(false, err.Error())
        return
    }
}

func (this *BuildingController) Update() {
    building := m.StoriedBuilding{}
    if err := this.ParseForm(&building); err != nil {
        this.Rsp(false, err.Error())
        return
    }
    id, err := m.UpdateBuilding(&building)
    if err == nil && id > 0 {
        this.Rsp(true, "Success")
        return
    } else {
        this.Rsp(false, err.Error())
        return
    }
}

func (this *BuildingController) GetAllBuildingList() {
    buildings, _ := m.GetBuildingList(1, 5000, "Id", m.BuildingQueryParam{})
    this.Data["json"] = buildings
    this.ServeJSON()
    return
}

func (this *BuildingController) GetUnitListByBuildingId() {
    Id, _ := this.GetInt64("BuildingId")
    houses, _ := m.GetHouseList(1, 5000, "UnitName", m.HouseQueryParam{BuildingId:Id})

    for _, house := range houses {
        if house.Owner != nil{
            house.Owner.User = nil
            house.Owner.Houses = nil
        }
        house.Building.Houses = nil
    }
    //简单去重
    found := make(map[string]bool)
    total := 0
    for i, val := range houses{
        if _, ok := found[val.UnitName]; !ok{
            found[val.UnitName] = true
            houses[total] = houses[i]
            total++
        }
    }
    houses = houses[:total]
    this.Data["json"] = houses
    this.ServeJSON()
    return
}
