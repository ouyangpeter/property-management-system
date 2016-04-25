package dao

import (
    . "os/user"
    "github.com/astaxie/beego/orm"
    . "property-management-system/models"
    "time"
)

var (
    GUserDao UserDao
)

func init() {
    GUserDao = NewUserDao()
}

type UserDao struct {
    ormer orm.Ormer
}

func NewUserDao() (userDao *UserDao) {
    userDao = &UserDao{}
    userDao.ormer = orm.NewOrm()
    return
}

func (this *UserDao) GetUserByUserId(uid uint64) (user *User, err error) {
    user = new(User)
    user.Uid = uid

    err = this.ormer.Read(user)
    return
}

func (this *UserDao) GetUserAuthByIdentityTypeAndIdentifier(identityType, identifier string) (userAuth UserAuth, err error) {
    userAuth = new(UserAuth)
    userAuth.IdentityType = identityType
    userAuth.Identifier = identifier
    err = this.ormer.Read(userAuth, "identity_type", "identifier")
    return
}

func (this *UserDao) UpdateUserAuth(userAuth UserAuth) (updateCnt int64, err error) {
    updateCnt, err = this.ormer.QueryTable("UserAuth").Filter("UserId", userAuth.User.UserId).Update(orm.Params{
        "Credential":userAuth.Credential,
        "Modified":time.Now(),
    })
    return
}