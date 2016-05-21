package pms

import (
    m "property-management-system/src/models"
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

func (this *BuildingController)AddBuilding() {
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

func (this *BuildingController) DeleteBuilding() {
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

func (this *BuildingController) UpdateBuilding() {
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