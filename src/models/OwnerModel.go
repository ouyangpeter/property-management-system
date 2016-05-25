package models

import (
    "github.com/astaxie/beego/orm"
    "time"
    "github.com/astaxie/beego/validation"
    "log"
    "errors"
    "property-management-system/src/lib"
)

var password0 = "000000"

type Owner struct {
    Id          int64
    Created     time.Time   `orm:"type(datetime);auto_now_add" form:"-"`
    Modified    time.Time   `orm:"type(datetime);auto_now;null" form:"-"`
    Remark      string      `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)"`
    Name        string      `orm:"size(32)" form:"Name"`
    PhoneNumber string      `orm:"size(32)" form:"PhoneNumber" valid:"Mobile"`
    IdCard      string      `orm:"size(32)" form:"IdCard" valid:"Length(18)"`
    Company     string      `orm:"size(32)" form:"Company"`
    Houses      []*House    `orm:"reverse(many)"`
    User        *User       `orm:"rel(one);on_delete(cascade)"`
    HouseId     string      `orm:"-" form:"HouseId" valid:"Required;"`
    UserName    string      `orm:"-" form:"UserName" valid:"Required;"`
}

func init() {
    orm.RegisterModel(new(Owner))
}

type OwnerQueryParam struct {
    //building_id:101
    //unit_name:2
    //house_id:1
    //owner_name:123
    //owner_phone:456
    //owner_idcard:789
    //owner_company:10
    Name        string
    PhoneNumber string
    IdCard      string
    Company     string
    BuildingId  int64
    UnitName    string
    HouseId     int64
}

func GetOwnerList(page int64, page_size int64, sort string, queryData OwnerQueryParam) (owners []Owner, count int64) {
    o := orm.NewOrm()
    qs := o.QueryTable(new(Owner))
    //todo need filter
    if queryData.BuildingId > 0 {
        qs = qs.Filter("Houses__Building__id", queryData.BuildingId)
    }
    if len(queryData.UnitName) > 0 {
        qs = qs.Filter("Houses__UnitName", queryData.UnitName)
    }
    if queryData.HouseId > 0 {
        qs = qs.Filter("Houses__Id", queryData.HouseId)
    }

    if len(queryData.Name) > 0 {
        qs = qs.Filter("Name__contains", queryData.Name)
    }
    if len(queryData.PhoneNumber) > 0{
        qs = qs.Filter("PhoneNumber__contains", queryData.PhoneNumber)
    }
    if len(queryData.IdCard) > 0{
        qs = qs.Filter("IdCard__contains", queryData.IdCard)
    }
    if len(queryData.Company) > 0{
        qs = qs.Filter("Company__contains", queryData.Company)
    }

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

func checkOwner(owner *Owner) (error) {
    valid := validation.Validation{}
    b, _ := valid.Valid(owner)
    if !b {
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
            return errors.New(err.Key + " " + err.Message)
        }
    }
    return nil
}
func AddOwner(owner *Owner) (int64, error) {
    if err := checkOwner(owner); err != nil {
        return 0, err
    }
    o := orm.NewOrm()
    terr := o.Begin()
    //先插入user
    user := new(User)
    user.UserName = owner.UserName
    user.Password = lib.Pwdhash(password0)
    user.Status = 2
    user.Type = 11
    _, err := o.Insert(user)

    if err != nil {
        terr = o.Rollback()
        return 0, err
    }
    //再插入owner
    newOwner := new(Owner)
    newOwner.Name = owner.Name
    newOwner.User = user
    newOwner.Company = owner.Company
    newOwner.Remark = owner.Remark
    newOwner.IdCard = owner.IdCard
    newOwner.PhoneNumber = owner.PhoneNumber
    id, err := o.Insert(newOwner)
    if err != nil {
        terr = o.Rollback()
        return 0, err
    }

    //如果有house,则更新house的owner
    house := new(House)
    err = o.QueryTable(new(House)).Filter("Id", owner.HouseId).One(house)
    if err != nil {
        terr = o.Rollback()
        return 0, err
    }
    if house.Owner != nil {
        terr = o.Rollback()
        return 0, errors.New("House already has owner")
    }
    house.Owner = newOwner
    _, err = o.Update(house)
    if err != nil {
        terr = o.Rollback()
        return 0, err
    } else {
        terr = o.Commit()
    }

    return id, terr

}

func DeleteOwnerById(Id int64) (int64, error) {
    o := orm.NewOrm()
    terr := o.Begin()
    owner := new(Owner)
    err := o.QueryTable(new(Owner)).Filter("Id", Id).RelatedSel().One(owner)
    if err != nil {
        terr = o.Rollback()
        return 0, err
    }
    if owner.User == nil {
        return 0, errors.New("Owner does not have user")
    }
    //把user也删了
    status, err := o.Delete(&Owner{Id:Id})
    if err == nil {
        status, err = o.Delete(owner.User)
    }
    if err != nil {
        terr = o.Rollback()
        return 0, err
    } else {
        terr = o.Commit()
    }
    return status, terr
}

func UpdateOwner(owner *Owner) (int64, error) {
    //formValid有点问题
    //if err := checkOwner(owner); err != nil {
    //    return 0, err
    //}
    o := orm.NewOrm()
    newOwner := make(orm.Params)
    if len(owner.PhoneNumber) > 0 {
        newOwner["PhoneNumber"] = owner.PhoneNumber
    }

    if len(owner.IdCard) > 0 {
        newOwner["IdCard"] = owner.IdCard
    }

    if len(owner.Company) > 0 {
        newOwner["Company"] = owner.Company
    }

    if len(owner.Remark) > 0 {
        newOwner["Remark"] = owner.Remark
    }
    if len(newOwner) == 0 {
        return 0, errors.New("update field is empty")
    }
    num, err := o.QueryTable(new(Owner)).Filter("Id", owner.Id).Update(newOwner)
    return num, err

}