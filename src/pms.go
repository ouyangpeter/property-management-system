package src

import (
    m "property-management-system/src/models"
    "errors"
    . "property-management-system/src/lib"
    "github.com/astaxie/beego"
    "log"
)

var PMS_AUTH_GATEWAY string = beego.AppConfig.String("pms_auth_gateway")

func CheckLogin(username string, password string) (user m.User, err error) {
    user = m.GetUserByUsername(username)
    if user.Id == 0 {
        return user, errors.New("用户不存在")
    }
    log.Println(user.Password)
    log.Println(Pwdhash(password))
    if user.Password != Pwdhash(password) {
        return user, errors.New("密码错误")
    }
    return user, nil
}
