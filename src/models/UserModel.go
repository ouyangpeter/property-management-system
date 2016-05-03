package models

import (
    "time"
    "github.com/astaxie/beego/orm"
)

type User struct {
    Id            int64
    UserName      string    `orm:"unique;size(32)" form:"Username"  valid:"Required;MaxSize(20);MinSize(6)"`
    Password      string    `orm:"size(32)" form:"Password" valid:"Required;MaxSize(40);MinSize(6)"`
    RePassword    string    `orm:"-" form:"Repassword" valid:"Required"`
    Email         string    `orm:"size(32)" form:"Email" valid:"Email"`
    Remark        string    `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)"`
    Status        int       `orm:"default(2)" form:"Status" valid:"Range(1,2)"`
    LastLoginTime time.Time `orm:"null;type(datetime)" form:"-"`
    Created       time.Time `orm:"type(datetime);auto_now_add"`
    Modified      time.Time `orm:"type(datetime):auto_now;null"`
    Roles         []*Role   `orm:"rel(m2m)"`
}

func init() {
    orm.RegisterModel(new(User))
}

func GetUserByUsername(userName string) (user User) {
    user = User{UserName: userName}
    o := orm.NewOrm()
    o.Read(&user, "UserName")
    return user
}

