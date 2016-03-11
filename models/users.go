package models

import (
    "github.com/astaxie/beego/orm"
    "time"
)

func init() {
    orm.RegisterModel(new(User))
    orm.RegisterModel(new(UserAuth))
}

type User struct {
    UserId    uint64 `orm:"auto; pk; not null"`
    Name      string `orm:"not null; size(30)"`
    Created   time.Time `orm:"auto_now; type(datetime)"`
    Modified  time.Time `orm:"auto_now_add; type(datetime)"`
    IsEnabled bool `orm:"not null; default(true)"`
}

type UserAuth struct {
    Id           uint64 `orm:"auto; pk; not null"`
    UserId       uint64 `orm:"not null"`
    IdentityType string `orm:"not null; size(30)", json:"IdentityType"`
    Identifier   string `orm:"not null; size(30)", json:"Identifier"`
    Credential   string `orm:"not null; size(50)", json:"Credential"`
}

func (this *UserAuth) IsUserPassword() bool {
    identityType := this.IdentityType
    if identityType == "username" || identityType == "phone" || identityType == "email" {
        return true
    }
    return false
}