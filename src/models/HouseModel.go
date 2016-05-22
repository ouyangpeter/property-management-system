package models

import "github.com/astaxie/beego/orm"

type House struct {
    Id       int64
    UnitName string
    Name     string
    Area     int
    Owner    *Owner `orm:"null;rel(fk);on_delete(set_null)"`
}

func init() {
    orm.RegisterModel(new(House))
}