package models

import (
    "time"
    "github.com/astaxie/beego/validation"
    "log"
    "errors"
    "github.com/astaxie/beego/orm"
)

type Notice struct {
    Id       int64
    Created  time.Time     `orm:"type(datetime);auto_now_add" form:"-"`
    Modified time.Time     `orm:"type(datetime);auto_now;null" form:"-"`
    Title    string        `orm:"size(32)" form:"Title" valid:"Required;"`
    Content  string        `orm:"size(1000)" form:"Content" valid:"Required;"`
}

func init() {
    orm.RegisterModel(new(Notice))
}

type NoticeQueryParam struct {
    Id int64
}

func checkNotice(notice *Notice) (error) {
    valid := validation.Validation{}
    h, _ := valid.Valid(&notice)
    if !h {
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
            return errors.New(err.Message)
        }
    }
    return nil
}

func GetNoticeList(page int64, page_size int64, sort string, queryData NoticeQueryParam) (notices []orm.Params, count int64) {
    o := orm.NewOrm()
    qs := o.QueryTable(new(Notice))
    if queryData.Id > 0 {
        qs = qs.Filter("Id", queryData.Id)
    }

    var offset int64
    if page <= 1 {
        offset = 0
    } else {
        offset = (page - 1 ) * page_size
    }
    qs.Limit(page_size, offset).OrderBy(sort).Values(&notices)
    count, _ = qs.Count()
    return notices, count
}

func GetNoticeById(id int64) (notice Notice, err error) {
    o := orm.NewOrm()
    err = o.QueryTable(new(Notice)).Filter("Id", id).One(&notice)
    return notice, err
}

func AddNotice(n *Notice) (int64, error) {
    if err := checkNotice(n); err != nil {
        return 0, err
    }
    o := orm.NewOrm()
    notice := new(Notice)
    notice.Title = n.Title
    notice.Content = n.Content
    id, err := o.Insert(notice)
    return id, err
}

func DeleteNoticeById(Id int64) (int64, error) {
    o := orm.NewOrm()
    status, err := o.Delete(&Notice{Id:Id})
    return status, err
}