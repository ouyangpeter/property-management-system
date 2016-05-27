package models

import (
    "time"
    "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego/validation"
    "log"
    "errors"
)

type ParkingSpot struct {
    Id              int64
    Created         time.Time     `orm:"type(datetime);auto_now_add" form:"-"`
    Modified        time.Time     `orm:"type(datetime);auto_now;null" form:"-"`
    Remark          string        `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)"`
    ParkingSpotNo   string        `orm:"size(32)" form:"ParkingSpotNo" valid:"MaxSize(20);MinSize(1)"`
    CarLicencePlate string        `orm:"size(32)" form:"CarLicencePlate" valid:"MaxSize(10);MinSize(1)"`
    CarName         string        `orm:"size(32)" form:"CarName" valid:"MaxSize(15);MinSize(1)"`
    CarColor        string        `orm:"size(32)" form:"CarColor" valid:"MaxSize(15);MinSize(1)"`
    Owner           *Owner        `orm:"null; rel(fk); on_delete(set_null)"`
    ParkingLot      *ParkingLot   `orm:"not null; rel(fk); on_delete(cascade)"`
    ParkingLotId    int64         `orm:"-" form:"ParkingLotId"`
    OwnerName       string        `orm:"-" form:"OwnerName"`
    OwnerPhone      string        `orm:"-" form:"OwnerPhone"`
}

type ParkingSpotQueryParam struct {
    ParkingLotId    int64
    ParkingSpotNo   string
    CarLicencePlate string
    CarName         string
    CarColor        string
    OwnerName       string
}

func init() {
    orm.RegisterModel(new(ParkingSpot))
}

func checkParkingSpot(parkingSpot *ParkingSpot) (error) {
    valid := validation.Validation{}
    h, _ := valid.Valid(&parkingSpot)
    if !h {
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
            return errors.New(err.Message)
        }
    }
    return nil
}

func GetParkingSpotList(page int64, page_size int64, sort string, queryData ParkingSpotQueryParam) (parkingSpots []ParkingSpot, count int64) {
    o := orm.NewOrm()
    qs := o.QueryTable(new(ParkingSpot))

    if queryData.ParkingLotId > 0 {
        qs = qs.Filter("ParkingLot__Id", queryData.ParkingLotId)
    }
    if len(queryData.ParkingSpotNo) > 0 {
        qs = qs.Filter("ParkingSpotNo__contains", queryData.ParkingSpotNo)
    }

    if len(queryData.CarLicencePlate) > 0 {
        qs = qs.Filter("CarLicencePlate__contains", queryData.CarLicencePlate)
    }

    if len(queryData.CarName) > 0 {
        qs = qs.Filter("CarName__contains", queryData.CarName)
    }
    if len(queryData.CarColor) > 0 {
        qs = qs.Filter("CarColor__contains", queryData.CarColor)
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
    qs.Limit(page_size, offset).OrderBy(sort).RelatedSel().All(&parkingSpots)
    count, _ = qs.Count()
    return parkingSpots, count
}

func AddParkingSpot(p *ParkingSpot) (int64, error) {
    log.Println(p)
    if err := checkParkingSpot(p); err != nil {
        return 0, err
    }
    o := orm.NewOrm()
    parkingLot := new(ParkingLot)
    parkingSpot := new(ParkingSpot)
    err := o.QueryTable(new(ParkingLot)).Filter("Id", p.ParkingLotId).One(parkingLot)
    if err != nil {
        return 0, err
    }
    parkingSpot.ParkingSpotNo = p.ParkingSpotNo
    parkingSpot.CarColor = p.CarColor
    parkingSpot.CarLicencePlate = p.CarLicencePlate
    parkingSpot.CarName = p.CarName
    parkingSpot.ParkingLot = parkingLot
    parkingSpot.Remark = p.Remark
    id, err := o.Insert(parkingSpot)
    return id, err
}

func DeleteParkingSpotById(Id int64) (int64, error) {
    o := orm.NewOrm()
    status, err := o.Delete(&ParkingSpot{Id:Id})
    return status, err
}

func UpdateParkingSpot(parkingSpot *ParkingSpot) (int64, error) {
    if err := checkParkingSpot(parkingSpot); err != nil {
        return 0, err
    }

    o := orm.NewOrm()
    terr := o.Begin()
    newParkingSpot := make(orm.Params)
    if len(parkingSpot.ParkingSpotNo) > 0 {
        newParkingSpot["ParkingSpotNo"] = parkingSpot.ParkingSpotNo
    }

    if len(parkingSpot.CarColor) > 0 {
        newParkingSpot["CarColor"] = parkingSpot.CarColor
    }

    if len(parkingSpot.CarLicencePlate) > 0 {
        newParkingSpot["CarLicencePlate"] = parkingSpot.CarLicencePlate
    }

    if len(parkingSpot.CarName) > 0 {
        newParkingSpot["CarName"] = parkingSpot.CarName
    }

    if len(parkingSpot.Remark) > 0 {
        newParkingSpot["Remark"] = parkingSpot.Remark
    }

    if len(parkingSpot.OwnerName) > 0 && len(parkingSpot.OwnerPhone) > 0 {
        var owner Owner
        err := o.QueryTable(new(Owner)).Filter("Name", parkingSpot.OwnerName).Filter("PhoneNumber", parkingSpot.OwnerPhone).One(&owner)
        if err != nil {
            return 0, err
        }
        var tParkingSpot ParkingSpot
        err = o.QueryTable(new(ParkingSpot)).Filter("Id", parkingSpot.Id).One(&tParkingSpot)
        if err != nil {
            return 0, err
        }

        tParkingSpot.Owner = &owner
        if _, err := o.Update(&tParkingSpot); err != nil {
            terr = o.Rollback()
            return 0, err
        }
    }

    if len(newParkingSpot) == 0 {
        return 0, errors.New("update field is empty")
    }

    num, err := o.QueryTable(new(ParkingSpot)).Filter("Id", parkingSpot.Id).Update(newParkingSpot)
    if err != nil {
        terr = o.Rollback()
        return 0, err
    }else{
        terr = o.Commit()
    }
    return num, terr
}

func RepealOwnerFromParkingSpot(Id int64) (int64 ,error){
    o := orm.NewOrm()
    var parkingSpot ParkingSpot
    err := o.QueryTable(new(ParkingSpot)).Filter("Id", Id).One(&parkingSpot)
    if err != nil {
        return 0, err
    }
    parkingSpot.CarColor = "";
    parkingSpot.CarName = "";
    parkingSpot.CarLicencePlate = "";
    parkingSpot.Owner = nil;
    num, err := o.Update(&parkingSpot)
    return num, err

}