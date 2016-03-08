package routers

import (
    "github.com/astaxie/beego"
    "property-management-system/controllers"
)

func init() {
    beego.Router("/login", &controllers.LoginController{})

}
