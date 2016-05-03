package src
import(
    m "property-management-system/src/models"
    "errors"
    . "property-management-system/src/lib"
)

func CheckLogin(username string, password string) (user m.User, err error) {
    user = m.GetUserByUsername(username)
    if user.Id == 0 {
        return user, errors.New("用户不存在")
    }
    if user.Password != Pwdhash(password) {
        return user, errors.New("密码错误")
    }
    return user, nil
}
