package models

import (
    "time"
    "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego/validation"
    "log"
    "errors"
)

type ParkingLot struct {
    Id             int64
    Created        time.Time     `orm:"type(datetime);auto_now_add" form:"-"`
    Modified       time.Time     `orm:"type(datetime);auto_now;null" form:"-"`
    Name           string        `orm:"unique;size(32)" form:"Name" valid:"Required;MaxSize(20);MinSize(1)"`
    Area           int           `form:"Area" valid:"Required"`
    ParkingSpotNum int           `form:"ParkingSpotNum" valid:"Required"`
    Remark         string        `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)"`
    ParkingSpots   []*ParkingSpot   `orm:"reverse(many)"`
}

func init() {
    orm.RegisterModel(new(ParkingLot))
}

type ParkingLotQueryParam struct {
    Name           string
    Area           int
    ParkingSpotNum int
}

func checkParkingLot(parkingLot *ParkingLot) (error) {
    valid := validation.Validation{}
    h, _ := valid.Valid(&parkingLot)
    if !h {
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
            return errors.New(err.Message)
        }
    }
    return nil
}

func GetParkingLotList(page int64, page_size int64, sort string, queryData ParkingLotQueryParam) (parkingLots []orm.Params, count int64) {
    o := orm.NewOrm()
    qs := o.QueryTable(new(ParkingLot))
    //log.Println(queryData)
    if len(queryData.Name) > 0 {
        qs = qs.Filter("Name__contains", queryData.Name)
    }
    if queryData.Area > 0 {
        qs = qs.Filter("Area", queryData.Area)
    }
    if queryData.ParkingSpotNum > 0 {
        qs = qs.Filter("ParkingSpotNum", queryData.ParkingSpotNum)
    }

    var offset int64
    if page <= 1 {
        offset = 0
    } else {
        offset = (page - 1 ) * page_size
    }
    qs.Limit(page_size, offset).OrderBy(sort).Values(&parkingLots)
    count, _ = qs.Count()
    return parkingLots, count
}

func AddParkingLot(p *ParkingLot) (int64, error) {
    if err := checkParkingLot(p); err != nil {
        return 0, err
    }
    o := orm.NewOrm()
    parkingLot := new(ParkingLot)
    parkingLot.Remark = p.Remark
    parkingLot.Area = p.Area
    parkingLot.Name = p.Name
    parkingLot.ParkingSpotNum = p.ParkingSpotNum
    id, err := o.Insert(parkingLot)
    return id, err
}

func DeleteParkingLotById(Id int64) (int64, error) {
    o := orm.NewOrm()
    status, err := o.Delete(&ParkingLot{Id:Id})
    return status, err
}

func UpdateParkingLot(parkingLot *ParkingLot) (int64, error) {
    if err := checkParkingLot(parkingLot); err != nil {
        return 0, err
    }
    o := orm.NewOrm()
    newParkingLot := make(orm.Params)

    if parkingLot.ParkingSpotNum != 0 {
        newParkingLot["ParkingSpotNum"] = parkingLot.ParkingSpotNum
    }

    if parkingLot.Area != 0 {
        newParkingLot["Area"] = parkingLot.Area
    }

    if len(parkingLot.Name) > 0 {
        newParkingLot["Name"] = parkingLot.Name
    }

    if len(parkingLot.Remark) > 0 {
        newParkingLot["Remark"] = parkingLot.Remark
    }

    if len(newParkingLot) == 0 {
        return 0, errors.New("update field is empty")
    }
    num, err := o.QueryTable(new(ParkingLot)).Filter("Id", parkingLot.Id).Update(newParkingLot)
    return num, err
}

