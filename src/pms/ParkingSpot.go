package pms

import (
    m "property-management-system/src/models"
)

type ParkingSpotController struct {
    CommonController
}

func (this *ParkingSpotController)Index() {
    page, _ := this.GetInt64("page")
    page_size, _ := this.GetInt64("rows")
    sort := this.GetString("sort")
    order := this.GetString("order")

    parkingSpotNo := this.GetString("parking_spot_no")
    parkingLotId, _ := this.GetInt64("parking_lot_id")
    ownerName := this.GetString("owner_name")

    carLicencePlate := this.GetString("car_licence_plate")
    carColor := this.GetString("car_color")
    carName := this.GetString("car_name")
    queryData := m.ParkingSpotQueryParam{
        ParkingLotId:parkingLotId,
        ParkingSpotNo:parkingSpotNo,
        OwnerName:ownerName,
        CarLicencePlate:carLicencePlate,
        CarColor:carColor,
        CarName:carName,
    }

    if len(order) > 0 {
        if order == "desc" {
            sort = "-" + sort
        }
    } else {
        sort = "Id"
    }
    parkingSpots, count := m.GetParkingSpotList(page, page_size, sort, queryData)

    for _, parkingSpot := range parkingSpots {
        if parkingSpot.Owner != nil {
            parkingSpot.Owner.User = nil
            parkingSpot.Owner.Houses = nil
            parkingSpot.Owner.ParkingSpots = nil
        }
        parkingSpot.ParkingLot.ParkingSpots = nil
    }

    if this.IsAjax() {
        if parkingSpots == nil {
            parkingSpots = make([]m.ParkingSpot, 0)
        }
        this.Data["json"] = &map[string]interface{}{"total":count, "rows":parkingSpots}
        //log.Println(houses[0].Building)
        this.ServeJSON()
        return
    } else {
        this.Data["parkingSpots"] = &parkingSpots
        if this.GetTemplateType() != "easyui" {
            this.Layout = this.GetTemplateType() + "/public/layout.tpl"
        }
        this.TplName = this.GetTemplateType() + "/pms/parkingSpot.tpl"
    }

}

func (this *ParkingSpotController)Add() {
    parkingSpot := m.ParkingSpot{}
    if err := this.ParseForm(&parkingSpot); err != nil {
        this.Rsp(false, err.Error())
        return
    }
    id, err := m.AddParkingSpot(&parkingSpot)
    if id > 0 && err == nil {
        this.Rsp(true, "Success")
        return
    } else {
        this.Rsp(false, err.Error())
        return
    }
}

func (this *ParkingSpotController)Delete() {
    Id, _ := this.GetInt64("Id")
    status, err := m.DeleteParkingSpotById(Id)
    if status > 0 && err == nil {
        this.Rsp(true, "Success")
        return
    } else {
        this.Rsp(false, err.Error())
        return
    }
}

func (this *ParkingSpotController)Update() {
    parkingSpot := m.ParkingSpot{}
    if err := this.ParseForm(&parkingSpot); err != nil {
        this.Rsp(false, err.Error())
        return
    }
    id, err := m.UpdateParkingSpot(&parkingSpot)
    if err == nil && id > 0 {
        this.Rsp(true, "Success")
        return
    } else {
        this.Rsp(false, err.Error())
        return
    }
}