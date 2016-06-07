package pms

import (
    m "property-management-system/src/models"
    "log"
)

type ComplaintController struct {
    CommonController
}

func (this *CommonController)Index() {
    page, _ := this.GetInt64("page")
    page_size, _ := this.GetInt64("rows")
    sort := this.GetString("sort")
    order := this.GetString("order")

    ownerName := this.GetString("owner_name")
    content := this.GetString("content")
    title := this.GetString("title")
    manager := this.GetString("Manager")
    status, _ := this.GetInt("status")

    queryData := m.ComplaintQueryParam{
        //Title     string
        //Content   string
        //OwnerName string
        //Manager   string
        //Status    int
        //Id        int64
        Content:content,
        Title:title,
        OwnerName:ownerName,
        Manager:manager,
        Status:status,
    }

    if len(order) > 0 {
        if order == "desc" {
            sort = "-" + sort
        }
    } else {
        sort = "Id"
    }
    userinfo := this.GetSession("userinfo")
    if userinfo == nil {
        this.Abort("401")
    }
    if (userinfo.(m.User).Type == 11) {
        owner, err := m.GetOwnerByUserId(userinfo.(m.User).Id)
        if err != nil {
            this.Abort("401")
        }
        queryData.OwnerId = owner.Id
    }
    complaints, count := m.GetComplaintList(page, page_size, sort, queryData)

    for _, complaint := range complaints {
        if complaint.Owner != nil {
            complaint.Owner.User = nil
            complaint.Owner.Houses = nil
            complaint.Owner.ParkingSpots = nil
            complaint.Owner.Complaints = nil
        }
    }

    if this.IsAjax() {
        if complaints == nil {
            complaints = make([]m.Complaint, 0)
        }
        this.Data["json"] = &map[string]interface{}{"total":count, "rows":complaints}
        //log.Println(houses[0].Building)
        this.ServeJSON()
        return
    } else {
        this.Data["userinfo"] = userinfo
        this.Data["complaints"] = &complaints
        if this.GetTemplateType() != "easyui" {
            this.Layout = this.GetTemplateType() + "/public/layout.tpl"
        }
        this.TplName = this.GetTemplateType() + "/pms/complaint.tpl"
    }

}

func (this *ComplaintController)Add() {
    complaint := m.Complaint{}
    if err := this.ParseForm(&complaint); err != nil {
        this.Rsp(false, err.Error())
        return
    }
    userinfo := this.GetSession("userinfo")
    if userinfo == nil {
        this.Abort("401")
    }
    if (userinfo.(m.User).Type == 11) {
        owner, err := m.GetOwnerByUserId(userinfo.(m.User).Id)
        if err != nil {
            this.Abort("401")
        }
        complaint.OwnerId = owner.Id
    }
    id, err := m.AddComplaint(&complaint)
    if id > 0 && err == nil {
        this.Rsp(true, "Success")
        return
    } else {
        this.Rsp(false, err.Error())
        return
    }
}

func (this *ComplaintController)Delete() {
    Id, _ := this.GetInt64("Id")
    status, err := m.DeleteComplaintById(Id)
    if status > 0 && err == nil {
        this.Rsp(true, "Success")
        return
    } else {
        this.Rsp(false, err.Error())
        return
    }
}

func (this *ComplaintController)Update() {
    complaint := m.Complaint{}
    if err := this.ParseForm(&complaint); err != nil {
        this.Rsp(false, err.Error())
        return
    }
    id, err := m.UpdateComplaint(&complaint)
    if err == nil && id > 0 {
        this.Rsp(true, "Success")
        return
    } else {
        this.Rsp(false, err.Error())
        return
    }

}

func (this *ComplaintController)Detail() {
    id, _ := this.GetInt64("Id")

    complaint, err := m.GetComplaintById(id)
    if err != nil {
        this.Abort("401")
    }
    if complaint.Owner != nil {
        complaint.Owner.User = nil
        complaint.Owner.Houses = nil
        complaint.Owner.ParkingSpots = nil
        complaint.Owner.Complaints = nil
    }
    if this.IsAjax() {
        this.Data["json"] = &complaint
        this.ServeJSON()
    } else {

        this.Data["complaint"] = &complaint
        log.Println(complaint)
        log.Println(complaint.Owner.Name)
        this.TplName = this.GetTemplateType() + "/pms/complaint_detail.tpl"
    }
}


