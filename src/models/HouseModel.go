package models

import (
    "github.com/astaxie/beego/orm"
    "time"
    "github.com/astaxie/beego/validation"
    "log"
    "errors"
)

type House struct {
    Id         int64
    Created    time.Time     `orm:"type(datetime);auto_now_add" form:"-"`
    Modified   time.Time     `orm:"type(datetime);auto_now;null" form:"-"`
    UnitName   string        `orm:"size(32)" form:"UnitName" valid:"Required;"`
    HouseNo    string        `orm:"size(32)" form:"HouseNo" valid:"Required;"`
    Area       int           `form:"Area" valid:"Required;"`
    Remark     string        `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)"`
    Owner      *Owner             `orm:"null; rel(fk); on_delete(set_null)"`
    Building   *StoriedBuilding   `orm:"not null; rel(fk); on_delete(cascade)"`
    BuildingId int64         `orm:"-" form:"BuildingId" valid:"Required;"`
}

type HouseQueryParam struct {
    UnitName   string
    HouseNo    string
    Area       int
    OwnerName  string
    BuildingId int64
}

func init() {
    orm.RegisterModel(new(House))
}

func checkHouse(house *House) (error) {
    valid := validation.Validation{}
    h, _ := valid.Valid(&house)
    if !h {
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
            return errors.New(err.Message)
        }
    }
    return nil
}

func GetHouseList(page int64, page_size int64, sort string, queryData HouseQueryParam) (houses []House, count int64) {
    o := orm.NewOrm()
    house := new(House)
    qs := o.QueryTable(house)
    if len(queryData.UnitName) > 0 {
        qs = qs.Filter("UnitName__contains", queryData.UnitName)
    }

    if len(queryData.HouseNo) > 0 {
        qs = qs.Filter("HouseNo__contains", queryData.HouseNo)
    }

    if queryData.Area > 0 {
        qs = qs.Filter("Area", queryData.Area)
    }

    if queryData.BuildingId > 0 {
        qs = qs.Filter("Building__Id", queryData.BuildingId)
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
    qs.Limit(page_size, offset).OrderBy(sort).RelatedSel().All(&houses)
    count, _ = qs.Count()
    return houses, count
}

func AddHouse(h *House) (int64, error) {
    if err := checkHouse(h); err != nil {
        return 0, err
    }
    o := orm.NewOrm()
    house := new(House)
    var buildingtable StoriedBuilding
    building := new(StoriedBuilding)
    err := o.QueryTable(buildingtable).Filter("Id", h.BuildingId).One(building)
    if err != nil {
        return 0, err
    }
    house.Building = building
    house.UnitName = h.UnitName
    house.HouseNo = h.HouseNo
    house.Area = h.Area
    house.Remark = h.Remark
    id, err := o.Insert(house)
    return id, err
}

func DeleteHouseById(Id int64) (int64, error) {
    o := orm.NewOrm()
    status, err := o.Delete(&House{Id:Id})
    return status, err
}

func UpdateHouse(house *House) (int64, error) {
    if err := checkHouse(house); err != nil {
        return 0, err
    }

    o := orm.NewOrm()
    newHouse := make(orm.Params)
    if len(house.UnitName) > 0 {
        newHouse["UnitName"] = house.UnitName
    }

    if len(house.HouseNo) > 0 {
        newHouse["HouseNo"] = house.HouseNo
    }

    if house.Area != 0 {
        newHouse["Area"] = house.Area
    }
    if len(house.Remark) > 0 {
        newHouse["Remark"] = house.Remark
    }

    if len(newHouse) == 0 {
        return 0, errors.New("update field is empty")
    }

    var houseTable House
    num, err := o.QueryTable(houseTable).Filter("Id", house.Id).Update(newHouse)
    return num, err
}

func GetHouseListByBuildingId(Id int64) (houses []House, count int64) {
    o := orm.NewOrm()
    house := new(House)
    qs := o.QueryTable(house)

    qs.OrderBy("Id").All(&houses)
    count, _ = qs.Count()
    return houses, count
}