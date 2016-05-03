package models

import (
    "time"
    "github.com/astaxie/beego/orm"
)

type Role struct {
    Id       int64
    Title    string  `orm:"size(100)" form:"Title"  valid:"Required"`
    Name     string  `orm:"size(100)" form:"Name"  valid:"Required"`
    Remark   string  `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)"`
    Status   int     `orm:"default(2)" form:"Status" valid:"Range(1,2)"`
    Users     []*User `orm:"reverse(many)"`
    Created  time.Time `orm:"type(datetime);auto_now_add"`
    Modified time.Time `orm:"type(datetime):auto_now;null"`
}

func init() {
    orm.RegisterModel(new(Role))
}
