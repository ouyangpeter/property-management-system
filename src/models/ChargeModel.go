package models

import (
    "time"
    "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego/validation"
    "log"
    "errors"
)

type Charge struct {
    Id           int64
    Created      time.Time     `orm:"type(datetime);auto_now_add" form:"-"`
    Modified     time.Time     `orm:"type(datetime);auto_now;null" form:"-"`
    Remark       string        `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)"`
    Money        int           `form:"Money"`
    //1 已结款 /  2 未结款
    Status       int           `orm:"default(2)" form:"Status" valid:"Range(1,2)"`
    CheckOutDate time.Time     `orm:"type(date);null" form:"CheckOutDate" valid:"Required"`
    ChargeType   *ChargeType   `orm:"not null; rel(fk); on_delete(do_nothing)"`
    Owner        *Owner        `orm:"null; rel(fk); on_delete(do_nothing)"`
    ChargeTypeId int64         `orm:"-" form:"ChargeTypeId"`
    OwnerId      int64         `orm:"-" form:"OwnerId"`
}

func init() {
    orm.RegisterModel(new(Charge))
}

type ChargeQueryParam struct {
    OwnerName    string
    ChargeTypeId int64
    Status       int
}

func checkCharge(charge *Charge) (error) {
    valid := validation.Validation{}
    h, _ := valid.Valid(&charge)
    if !h {
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
            return errors.New(err.Message)
        }
    }
    return nil
}

func GetChargeList(page int64, page_size int64, sort string, queryData ChargeQueryParam) (charges []Charge, count int64) {
    o := orm.NewOrm()
    qs := o.QueryTable(new(Charge))

    if queryData.ChargeTypeId > 0 {
        qs = qs.Filter("ChargeType__Id", queryData.ChargeTypeId)
    }

    if queryData.Status > 0 {
        qs = qs.Filter("Status", queryData.Status)
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
    qs.Limit(page_size, offset).OrderBy(sort).RelatedSel().All(&charges)
    count, _ = qs.Count()
    return charges, count
}

func AddCharge(c *Charge) (int64, error) {
    if err := checkCharge(c); err != nil {
        return 0, err
    }
    o := orm.NewOrm()
    chargeType := new(ChargeType)
    charge := new(Charge)
    err := o.QueryTable(new(ChargeType)).Filter("Id", c.ChargeTypeId).One(chargeType)
    if err != nil {
        return 0, err
    }
    owner := new(Owner)
    err = o.QueryTable(new(Owner)).Filter("Id", c.OwnerId).One(owner)
    if err != nil {
        return 0, err
    }

    charge.ChargeType = chargeType
    charge.Owner = owner
    charge.Status = 2
    charge.Remark = c.Remark
    charge.Money = c.Money

    id, err := o.Insert(charge)
    return id, err
}

func DeleteChargeById(Id int64) (int64, error) {
    o := orm.NewOrm()
    status, err := o.Delete(&Charge{Id:Id})
    return status, err
}

func UpdateCharge(charge *Charge) (int64, error) {
    if err := checkCharge(charge); err != nil {
        return 0, err
    }

    o := orm.NewOrm()
    newCharge := make(orm.Params)
    if charge.Money > 0 {
        newCharge["Money"] = charge.Money
    }

    if charge.Status > 0 {
        newCharge["Status"] = charge.Status
        if charge.Status == 1 {
            newCharge["CheckOutDate"] = time.Now()
        }
    }

    if len(charge.Remark) > 0 {
        newCharge["Remark"] = charge.Remark
    }

    if len(newCharge) == 0 {
        return 0, errors.New("update field is empty")
    }

    num, err := o.QueryTable(new(Charge)).Filter("Id", charge.Id).Update(newCharge)
    return num, err
}
