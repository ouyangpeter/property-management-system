package models

import "github.com/astaxie/beego/orm"

type Owner struct {
    Id          int64
    Name        string
    PhoneNumber string
    IdCard      string
    Houses      []*House `orm:"reverse(many)"`
}

func init() {
    orm.RegisterModel(new(Owner))
}