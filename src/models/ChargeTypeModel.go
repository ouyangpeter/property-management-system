package models

import (
    "time"
    "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego/validation"
    "log"
    "errors"
)

type ChargeType struct {
    Id        int64
    Created   time.Time     `orm:"type(datetime);auto_now_add" form:"-"`
    Modified  time.Time     `orm:"type(datetime);auto_now;null" form:"-"`
    Remark    string        `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)"`
    Name      string        `orm:"unique;size(32)" form:"Name" valid:"Required;MaxSize(20);MinSize(1)"`
    Criterion string        `orm:"size(32)" form:"Criterion" valid:"Required;"`
}

type ChargeTypeQueryParam struct {
    Name      string
    Criterion string
}

func init() {
    orm.RegisterModel(new(ChargeType))
}

func checkChargeType(chargeType *ChargeType) (error) {
    valid := validation.Validation{}
    h, _ := valid.Valid(&chargeType)
    if !h {
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
            return errors.New(err.Message)
        }
    }
    return nil
}

func GetChargeTypeList(page int64, page_size int64, sort string, queryData ChargeTypeQueryParam) (chargeTypes []orm.Params, count int64) {
    o := orm.NewOrm()
    qs := o.QueryTable(new(ChargeType))
    //log.Println(queryData)
    if len(queryData.Name) > 0 {
        qs = qs.Filter("Name__contains", queryData.Name)
    }
    if len(queryData.Criterion) > 0 {
        qs = qs.Filter("Criterion__contains", queryData.Criterion)
    }

    var offset int64
    if page <= 1 {
        offset = 0
    } else {
        offset = (page - 1 ) * page_size
    }
    qs.Limit(page_size, offset).OrderBy(sort).Values(&chargeTypes)
    count, _ = qs.Count()
    return chargeTypes, count
}

func AddChargeType(c *ChargeType) (int64, error) {
    if err := checkChargeType(c); err != nil {
        return 0, err
    }
    o := orm.NewOrm()
    chargeType := new(ChargeType)
    chargeType.Name = c.Name
    chargeType.Criterion = c.Criterion
    chargeType.Remark = c.Remark

    id, err := o.Insert(chargeType)
    return id, err
}

func DeleteChargeTypeById(Id int64) (int64, error) {
    o := orm.NewOrm()
    status, err := o.Delete(&ChargeType{Id:Id})
    return status, err
}

func UpdateChargeType(chargeType *ChargeType) (int64, error) {
    if err := checkChargeType(chargeType); err != nil {
        return 0, err
    }
    o := orm.NewOrm()
    newChargeType := make(orm.Params)

    if len(chargeType.Remark) > 0 {
        newChargeType["Remark"] = chargeType.Remark
    }

    if len(chargeType.Criterion) > 0 {
        newChargeType["Criterion"] = chargeType.Criterion
    }

    if len(newChargeType) == 0 {
        return 0, errors.New("update field is empty")
    }
    num, err := o.QueryTable(new(ChargeType)).Filter("Id", chargeType.Id).Update(newChargeType)
    return num, err
}

