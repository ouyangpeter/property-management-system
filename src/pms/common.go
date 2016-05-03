package pms

import (
    "github.com/astaxie/beego"
)

type CommonController struct {
    beego.Controller
}

func (this *CommonController)GetTemplateType() string {
    templateType := beego.AppConfig.String("template_type")
    if "" == templateType{
        templateType = "easyui"
    }
    return templateType
}

func (this *CommonController) Rsp(status bool, str string) {
    this.Data["json"] = &map[string]interface{}{"status": status, "info": str}
    this.ServeJSON()
}

func init() {

}
