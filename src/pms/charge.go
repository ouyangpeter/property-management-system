package pms

import (
    m "property-management-system/src/models"
)

type ChargeController struct {
    CommonController
}

func (this *ChargeController)Index() {
    page, _ := this.GetInt64("page")
    page_size, _ := this.GetInt64("rows")
    sort := this.GetString("sort")
    order := this.GetString("order")

    ownerName := this.GetString("owner_name")
    chargeTypeId, _ := this.GetInt64("charge_type_id")
    status, _ := this.GetInt("status")
    queryData := m.ChargeQueryParam{
        OwnerName:ownerName,
        ChargeTypeId:chargeTypeId,
        Status:status,
    }

    if len(order) > 0 {
        if order == "desc" {
            sort = "-" + sort
        }
    } else {
        sort = "Id"
    }
    charges, count := m.GetChargeList(page, page_size, sort, queryData)

    for _, charge := range charges {
        if charge.Owner != nil {
            charge.Owner.User = nil
            charge.Owner.Houses = nil
            charge.Owner.ParkingSpots = nil
            charge.Owner.Charges = nil
        }
        charge.ChargeType.Charges = nil
    }

    if this.IsAjax() {
        if charges == nil {
            charges = make([]m.Charge, 0)
        }
        this.Data["json"] = &map[string]interface{}{"total":count, "rows":charges}
        //log.Println(houses[0].Building)
        this.ServeJSON()
        return
    } else {
        this.Data["charges"] = &charges
        if this.GetTemplateType() != "easyui" {
            this.Layout = this.GetTemplateType() + "/public/layout.tpl"
        }
        this.TplName = this.GetTemplateType() + "/pms/charge.tpl"
    }

}

func (this *ChargeController)Add() {
    charge := m.Charge{}
    if err := this.ParseForm(&charge); err != nil {
        this.Rsp(false, err.Error())
        return
    }
    id, err := m.AddCharge(&charge)
    if id > 0 && err == nil {
        this.Rsp(true, "Success")
        return
    } else {
        this.Rsp(false, err.Error())
        return
    }
}

func (this *ChargeController)Delete() {
    Id, _ := this.GetInt64("Id")
    status, err := m.DeleteChargeById(Id)
    if status > 0 && err == nil {
        this.Rsp(true, "Success")
        return
    } else {
        this.Rsp(false, err.Error())
        return
    }
}

func (this *ChargeController)Update() {
    charge := m.Charge{}
    if err := this.ParseForm(&charge); err != nil {
        this.Rsp(false, err.Error())
        return
    }
    id, err := m.UpdateCharge(&charge)
    if err == nil && id > 0 {
        this.Rsp(true, "Success")
        return
    } else {
        this.Rsp(false, err.Error())
        return
    }
}

