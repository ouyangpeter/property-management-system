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
    UserAuthes []*UserAuth `orm:"reverse(many)"`
}

type UserAuth struct {
    Id           uint64 `orm:"auto; pk; not null"`
    User         *User `orm:"rel(fk)"`
    IdentityType string `orm:"not null; size(30)"`
    Identifier   string `orm:"not null; size(30)"`
    Credential   string `orm:"not null; size(50)"`
    Created      time.Time `orm:"auto_now; type(datetime)"`
    Modified     time.Time `orm:"auto_now_add; type(datetime)"`
}
