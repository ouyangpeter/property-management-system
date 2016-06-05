package models

import (
    "time"
    "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego/validation"
    "log"
    "errors"
)

type Complaint struct {
    Id          int64
    Created     time.Time     `orm:"type(datetime);auto_now_add" form:"-"`
    Modified    time.Time     `orm:"type(datetime);auto_now;null" form:"-"`
    Title       string        `orm:"size(32)" form:"Title" valid:"Required;"`
    Content     string        `orm:"size(1000)" form:"Content" valid:"Required;"`
    // 1 待处理 2已处理
    Status      int           `form:"Status" valid:"Required;"`
    Manager     string        `orm:"null; size(32)" form:"Manager"`
    Description string        `orm:"null; size(1000)" form:"Description"`
    Owner       *Owner        `orm:"null; rel(fk); on_delete(do_nothing)"`
    OwnerId     int64         `orm:"-" form:"OwnerId"`
}

type ComplaintQueryParam struct {
    Title     string
    Content   string
    OwnerName string
    Manager   string
    Status    int
    OwnerId   int64
}

func init() {
    orm.RegisterModel(new(Complaint))
}

func checkComplaint(complaint *Complaint) (error) {
    valid := validation.Validation{}
    h, _ := valid.Valid(&complaint)
    if !h {
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
            return errors.New(err.Message)
        }
    }
    return nil
}

func GetComplaintList(page int64, page_size int64, sort string, queryData ComplaintQueryParam) (complaints []Complaint, count int64) {
    o := orm.NewOrm()
    complaint := new(Complaint)
    qs := o.QueryTable(complaint)
    if len(queryData.Title) > 0 {
        qs = qs.Filter("Title__contains", queryData.Title)
    }

    if len(queryData.Content) > 0 {
        qs = qs.Filter("Content__contains", queryData.Content)
    }
    if len(queryData.Manager) > 0 {
        qs = qs.Filter("Manager__contains", queryData.Manager)
    }

    if queryData.Status > 0 {
        qs = qs.Filter("Status", queryData.Status)
    }

    if queryData.OwnerId > 0 {
        qs = qs.Filter("Owner__Id", queryData.OwnerId)
    }

    if len(queryData.OwnerName) > 0 {

        qs = qs.Filter("Owner__Name", queryData.OwnerName)
    }

    var offset int64
    if page <= 1 {
        offset = 0
    } else {
        offset = (page - 1) * page_size
    }
    qs.Limit(page_size, offset).OrderBy(sort).RelatedSel().All(&complaints)
    count, _ = qs.Count()
    return complaints, count
}

func AddComplaint(c *Complaint) (int64, error) {
    if err := checkComplaint(c); err != nil {
        return 0, err
    }
    o := orm.NewOrm()
    complaint := new(Complaint)
    owner := new(Owner)
    err := o.QueryTable(new(Owner)).Filter("Id", c.OwnerId).One(owner)
    if err != nil {
        return 0, err
    }
    complaint.Content = c.Content
    complaint.Title = c.Title
    complaint.Owner = owner
    complaint.Status = 1
    id, err := o.Insert(complaint)
    return id, err
}

func DeleteComplaintById(Id int64) (int64, error) {
    o := orm.NewOrm()
    status, err := o.Delete(&Complaint{Id:Id})
    return status, err
}

func UpdateComplaint(complaint *Complaint) (int64, error) {
    if err := checkComplaint(complaint); err != nil {
        return 0, err
    }

    o := orm.NewOrm()
    newComplaint := make(orm.Params)
    if len(complaint.Manager) > 0 {
        newComplaint["Manager"] = complaint.Manager
    }

    if len(complaint.Description) > 0 {
        newComplaint["Description"] = complaint.Description
    }

    if complaint.Status > 0 {
        newComplaint["Status"] = complaint.Status
    }

    if len(newComplaint) == 0 {
        return 0, errors.New("update field is empty")
    }

    num, err := o.QueryTable(new(Complaint)).Filter("Id", complaint.Id).Update(newComplaint)
    return num, err
}

func GetComplaintById(id int64)(complaint Complaint, err error){
    o := orm.NewOrm()
    err = o.QueryTable(new(Complaint)).Filter("Id", id).RelatedSel().One(&complaint)
    return complaint, err
}