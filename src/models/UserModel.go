package models

import (
    "github.com/astaxie/beego/orm"
    "time"
    "github.com/astaxie/beego/validation"
    "log"
    "errors"
    "property-management-system/src/lib"
)

type User struct {
    Id            int64
    UserName      string    `orm:"unique;size(32)" form:"Username"  valid:"Required;MaxSize(20);MinSize(6)"`
    Password      string    `orm:"size(128)" form:"Password" valid:"Required;MaxSize(70);MinSize(50)"`
    RePassword    string    `orm:"-" form:"Repassword" valid:"Required"`
    Nickname      string    `orm:"size(32)" form:"Nickname" valid:"Required;MaxSize(20);MinSize(2)"`
    Type          int       `orm:"default(11)" valid:"Required;Range(11,12)"`
    Email         string    `orm:"size(32)" form:"Email" valid:"Email"`
    Remark        string    `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)"`
    Status        int       `orm:"default(2)" form:"Status" valid:"Range(1,2)"`
    LastLoginTime time.Time `orm:"null;type(datetime)" form:"-"`
    Created       time.Time `orm:"type(datetime);auto_now_add" form:"-"`
    Modified      time.Time `orm:"type(datetime);auto_now;null" form:"-"`
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

func UpdateUser(u *User) (int64, error) {
    //log.Println(u)
    if err := checkUser(u); err != nil {
        return 0, err
    }
    o := orm.NewOrm()
    user := make(orm.Params)
    if len(u.UserName) > 0 {
        user["UserName"] = u.UserName
    }
    if len(u.Nickname) > 0 {
        user["Nickname"] = u.Nickname
    }
    if len(u.Email) > 0 {
        user["Email"] = u.Email
    }
    if len(u.Remark) > 0 {
        user["Remark"] = u.Remark
    }
    if len(u.Password) > 0 {
        user["Password"] = lib.Pwdhash(u.Password)
    }
    if u.Status != 0{
        user["Status"] = u.Status
    }
    if len(user) == 0{
        return 0, errors.New("update field is empty")
    }

    var table User
    num, err := o.QueryTable(table).Filter("Id", u.Id).Update(user)
    return num, err

}

func checkUser(u *User) (err error) {
    valid := validation.Validation{}
    b, _ := valid.Valid(&u)
    if !b {
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
            return errors.New(err.Message)
        }
    }
    return nil
}