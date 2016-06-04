package pms

import (
    m "property-management-system/src/models"
    "github.com/astaxie/beego/orm"
)

type NoticeController struct {
    CommonController
}

func (this *NoticeController)Index() {
    page, _ := this.GetInt64("page")
    page_size, _ := this.GetInt64("rows")
    sort := this.GetString("sort")
    order := this.GetString("order")
    id, _ := this.GetInt64("Id")
    queryData := m.NoticeQueryParam{
        Id:id,
    }

    if len(order) > 0 {
        if order == "desc" {
            sort = "-" + sort
        }
    } else {
        sort = "Id"
    }

    notices, count := m.GetNoticeList(page, page_size, sort, queryData)

    if this.IsAjax() {
        if notices == nil {
            notices = make([]orm.Params, 0)
        }
        this.Data["json"] = &map[string]interface{}{"total":count, "rows":notices}
        this.ServeJSON()
        return
    } else {
        this.Data["notices "] = &notices
        if this.GetTemplateType() != "easyui" {
            this.Layout = this.GetTemplateType() + "/public/layout.tpl"
        }
        this.TplName = this.GetTemplateType() + "/pms/notice.tpl"
    }
}

func (this *NoticeController)Add() {
    notice := m.Notice{}
    if err := this.ParseForm(&notice); err != nil {
        this.Rsp(false, err.Error())
        return
    }
    id, err := m.AddNotice(&notice)
    if id > 0 && err == nil {
        this.Rsp(true, "Success")
        return
    } else {
        this.Rsp(false, err.Error())
        return
    }
}

func (this *NoticeController) Delete() {
    Id, _ := this.GetInt64("Id")
    status, err := m.DeleteNoticeById(Id)
    if status > 0 && err == nil {
        this.Rsp(true, "Success")
        return
    } else {
        this.Rsp(false, err.Error())
        return
    }
}
