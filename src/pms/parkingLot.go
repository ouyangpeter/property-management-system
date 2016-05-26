package pms

import (
    m "property-management-system/src/models"
    "github.com/astaxie/beego/orm"
)

type ParkingLotController struct {
    CommonController
}

func (this *ParkingLotController)Index() {
    page, _ := this.GetInt64("page")
    page_size, _ := this.GetInt64("rows")
    sort := this.GetString("sort")
    order := this.GetString("order")

    name := this.GetString("name")
    parkingSpotNum, _ := this.GetInt("parking_spot_num")
    area, _ := this.GetInt("area")
    queryData := m.ParkingLotQueryParam{
        Name:name,
        ParkingSpotNum:parkingSpotNum,
        Area: area,
    }

    if len(order) > 0 {
        if order == "desc" {
            sort = "-" + sort
        }
    } else {
        sort = "Id"
    }
    parkingLots, count := m.GetParkingLotList(page, page_size, sort, queryData)

    if this.IsAjax() {
        if parkingLots == nil {
            parkingLots = make([]orm.Params, 0)
        }
        this.Data["json"] = &map[string]interface{}{"total":count, "rows":parkingLots}
        this.ServeJSON()
        return
    } else {
        this.Data["parkingLots"] = &parkingLots
        if this.GetTemplateType() != "easyui" {
            this.Layout = this.GetTemplateType() + "/public/layout.tpl"
        }
        this.TplName = this.GetTemplateType() + "/pms/parkingLot.tpl"
    }
}

func (this *ParkingLotController)Add() {
    parkingLot := m.ParkingLot{}
    if err := this.ParseForm(&parkingLot); err != nil {
        this.Rsp(false, err.Error())
        return
    }
    id, err := m.AddParkingLot(&parkingLot)
    if id > 0 && err == nil {
        this.Rsp(true, "Success")
        return
    } else {
        this.Rsp(false, err.Error())
        return
    }
}

func (this *ParkingLotController) Delete() {
    Id, _ := this.GetInt64("Id")
    status, err := m.DeleteParkingLotById(Id)
    if status > 0 && err == nil {
        this.Rsp(true, "Success")
        return
    } else {
        this.Rsp(false, err.Error())
        return
    }
}

func (this *ParkingLotController) Update() {
    parkingLot := m.ParkingLot{}
    if err := this.ParseForm(&parkingLot); err != nil {
        this.Rsp(false, err.Error())
        return
    }
    id, err := m.UpdateParkingLot(&parkingLot)
    if err == nil && id > 0 {
        this.Rsp(true, "Success")
        return
    } else {
        this.Rsp(false, err.Error())
        return
    }
}
