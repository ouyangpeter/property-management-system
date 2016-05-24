package models

import (
    "github.com/astaxie/beego/orm"
    "time"
)

type Owner struct {
    Id          int64
    Created     time.Time   `orm:"type(datetime);auto_now_add" form:"-"`
    Modified    time.Time   `orm:"type(datetime);auto_now;null" form:"-"`
    Remark        string    `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)"`
    Name        string      `orm:"size(32)"`
    PhoneNumber string      `orm:"size(32)"`
    IdCard      string      `orm:"size(32)"`
    Company     string      `orm:"size(32)"`
    Houses      []*House    `orm:"reverse(many)"`
    User        *User       `orm:"rel(one)"`
}

func init() {
    orm.RegisterModel(new(Owner))
}

type OwnerQueryParam struct {
    Name        string
    PhoneNumber string
    IdCard      string
    Company     string
    BuildingId  int64
    UnitName    string
}

func GetOwnerList(page int64, page_size int64, sort string, queryData OwnerQueryParam) (owners []Owner, count int64) {
    o := orm.NewOrm()
    qs := o.QueryTable(new(Owner))
    //todo need filter


    var offset int64
    if page <= 1 {
        offset = 0
    } else {
        offset = (page - 1) * page_size
    }
    qs.Limit(page_size, offset).OrderBy(sort).RelatedSel().All(&owners)
    count, _ = qs.Count()
    return owners, count
}
