package pms

import
(
    "github.com/astaxie/beego/orm"
    m "property-management-system/src/models"
)

type ChargeTypeController struct {
    CommonController
}

func (this *ChargeTypeController)Index() {
    page, _ := this.GetInt64("page")
    page_size, _ := this.GetInt64("rows")
    sort := this.GetString("sort")
    order := this.GetString("order")

    name := this.GetString("name")
    criterion := this.GetString("criterion")

    queryData := m.ChargeTypeQueryParam{
        Name:name,
        Criterion:criterion,
    }

    if len(order) > 0 {
        if order == "desc" {
            sort = "-" + sort
        }
    } else {
        sort = "Id"
    }
    chargeTypes, count := m.GetChargeTypeList(page, page_size, sort, queryData)

    if this.IsAjax() {
        if chargeTypes == nil {
            chargeTypes = make([]orm.Params, 0)
        }
        this.Data["json"] = &map[string]interface{}{"total":count, "rows":chargeTypes}
        this.ServeJSON()
        return
    } else {
        userinfo := this.GetSession("userinfo")
        if userinfo == nil {
            this.Abort("401")
        }
        this.Data["userinfo"] = userinfo
        this.Data["chargeTypes"] = &chargeTypes
        if this.GetTemplateType() != "easyui" {
            this.Layout = this.GetTemplateType() + "/public/layout.tpl"
        }
        this.TplName = this.GetTemplateType() + "/pms/chargeType.tpl"
    }
}

func (this *ChargeTypeController)Add() {
    chargeType := m.ChargeType{}
    if err := this.ParseForm(&chargeType); err != nil {
        this.Rsp(false, err.Error())
        return
    }
    id, err := m.AddChargeType(&chargeType)
    if id > 0 && err == nil {
        this.Rsp(true, "Success")
        return
    } else {
        this.Rsp(false, err.Error())
        return
    }
}

func (this *ChargeTypeController) Delete() {
    Id, _ := this.GetInt64("Id")
    status, err := m.DeleteChargeTypeById(Id)
    if status > 0 && err == nil {
        this.Rsp(true, "Success")
        return
    } else {
        this.Rsp(false, err.Error())
        return
    }
}

func (this *ChargeTypeController) Update() {
    chargeType := m.ChargeType{}
    if err := this.ParseForm(&chargeType); err != nil {
        this.Rsp(false, err.Error())
        return
    }
    id, err := m.UpdateChargeType(&chargeType)
    if err == nil && id > 0 {
        this.Rsp(true, "Success")
        return
    } else {
        this.Rsp(false, err.Error())
        return
    }
}

func (this *ChargeTypeController) GetAllChargeTypeList() {
    chargeTypes, _ := m.GetChargeTypeList(1, 5000, "Id", m.ChargeTypeQueryParam{})
    this.Data["json"] = chargeTypes
    this.ServeJSON()
    return
}
