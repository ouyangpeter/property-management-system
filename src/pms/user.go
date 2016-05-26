package pms

import (
    m "property-management-system/src/models"
    "github.com/astaxie/beego/orm"
)

type UserController struct {
    CommonController
}

func (this *UserController)Index() {
    page, _ := this.GetInt64("page")
    page_size, _ := this.GetInt64("rows")
    sort := this.GetString("sort")
    order := this.GetString("order")

    queryData := m.UserQueryParam{

    }

    if len(order) > 0 {
        if order == "desc" {
            sort = "-" + sort
        }
    } else {
        sort = "Id"
    }
    users, count := m.GetUserList(page, page_size, sort, queryData)

    if this.IsAjax() {
        if users == nil {
            users = make([]orm.Params, 0)
        }
        this.Data["json"] = &map[string]interface{}{"total":count, "rows":users}
        //log.Println(houses[0].Building)
        this.ServeJSON()
        return
    } else {
        this.Data["users"] = &users
        if this.GetTemplateType() != "easyui" {
            this.Layout = this.GetTemplateType() + "/public/layout.tpl"
        }
        this.TplName = this.GetTemplateType() + "/pms/user.tpl"
    }

}

