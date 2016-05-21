package models

import (
    "database/sql"
    "fmt"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "mime"
    "os"
    . "property-management-system/src/lib"
    "time"
    "strconv"
)

var o orm.Ormer

func Initialize() {
    mime.AddExtensionType(".css", "text/css")
    initArgs()
    Connect()
    beego.AddFuncMap("stringsToJson", StringsToJson)
}

func initArgs() {
    args := os.Args
    for _, v := range args {
        if v == "--syncdb" {
            Syncdb()
            os.Exit(0)
        }
    }
}

func Syncdb() {
    createdb()
    Connect()
    o = orm.NewOrm()
    // 数据库别名
    name := "default"
    // drop table 后再建表
    force := true
    // 打印执行过程
    verbose := true
    // 遇到错误立即返回
    err := orm.RunSyncdb(name, force, verbose)
    if err != nil {
        fmt.Println(err)
    }
    insertUser()
    for i := 0; i < 100; i++ {
        insertBuilding("test" + strconv.Itoa(i))
    }
    fmt.Println("database init is complete.\nPlease restart the application")

}

//数据库连接
func Connect() {
    var dns string
    db_type := beego.AppConfig.String("db_type")
    db_host := beego.AppConfig.String("db_host")
    db_port := beego.AppConfig.String("db_port")
    db_user := beego.AppConfig.String("db_user")
    db_pass := beego.AppConfig.String("db_pass")
    db_name := beego.AppConfig.String("db_name")
    //db_path := beego.AppConfig.String("db_path")
    //db_sslmode := beego.AppConfig.String("db_sslmode")
    switch db_type {
    case "mysql":
        orm.RegisterDriver("mysql", orm.DRMySQL)
        dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", db_user, db_pass, db_host, db_port, db_name)
        break
    //case "postgres":
    //    orm.RegisterDriver("postgres", orm.DRPostgres)
    //    dns = fmt.Sprintf("dbname=%s host=%s  user=%s  password=%s  port=%s  sslmode=%s", db_name, db_host, db_user, db_pass, db_port, db_sslmode)
    //case "sqlite3":
    //    orm.RegisterDriver("sqlite3", orm.DRSqlite)
    //    if db_path == "" {
    //        db_path = "./"
    //    }
    //    dns = fmt.Sprintf("%s%s.db", db_path, db_name)
    //    break
    default:
        beego.Critical("Database driver is not allowed:", db_type)
    }
    err := orm.RegisterDataBase("default", db_type, dns)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(-1)
    }
}

//创建数据库
func createdb() {

    db_type := beego.AppConfig.String("db_type")
    db_host := beego.AppConfig.String("db_host")
    db_port := beego.AppConfig.String("db_port")
    db_user := beego.AppConfig.String("db_user")
    db_pass := beego.AppConfig.String("db_pass")
    db_name := beego.AppConfig.String("db_name")
    //db_path := beego.AppConfig.String("db_path")
    //db_sslmode := beego.AppConfig.String("db_sslmode")

    var dns string
    var sqlstring string
    switch db_type {
    case "mysql":
        dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8", db_user, db_pass, db_host, db_port)
        sqlstring = fmt.Sprintf("CREATE DATABASE  if not exists `%s` CHARSET utf8 COLLATE utf8_general_ci", db_name)
        break
    //case "postgres":
    //    dns = fmt.Sprintf("host=%s  user=%s  password=%s  port=%s  sslmode=%s", db_host, db_user, db_pass, db_port, db_sslmode)
    //    sqlstring = fmt.Sprintf("CREATE DATABASE %s", db_name)
    //    break
    //case "sqlite3":
    //    if db_path == "" {
    //        db_path = "./"
    //    }
    //    dns = fmt.Sprintf("%s%s.db", db_path, db_name)
    //    os.Remove(dns)
    //    sqlstring = "create table init (n varchar(32));drop table init;"
    //    break
    default:
        beego.Critical("Database driver is not allowed:", db_type)
    }
    db, err := sql.Open(db_type, dns)
    if err != nil {
        panic(err.Error())
    }
    r, err := db.Exec(sqlstring)
    if err != nil {
        log.Println(err)
        log.Println(r)
    } else {
        log.Println("Database ", db_name, " created")
    }
    defer db.Close()

}

func insertUser() {
    fmt.Println("insert user ...")
    u := new(User)
    u.UserName = "admin"
    u.Password = Pwdhash("admin")
    u.Email = "ouyangpeter911@gmail.com"
    u.Remark = "I'm admin"
    u.Nickname = "欧阳"
    u.Status = 2
    u.Type = 12
    o = orm.NewOrm()
    n, err := o.Insert(u)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(-1)
    } else {
        fmt.Println("insert", n, "user end")
    }
}

func insertBuilding(name string) {
    fmt.Println("insert building:", name)
    b := new(StoriedBuilding)
    b.Name = name
    b.Height = 50
    b.Floors = 11
    b.Area = 800
    b.BuildDate = time.Now()
    o = orm.NewOrm()
    n, err := o.Insert(b)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(-1)
    } else {
        fmt.Println("insert", n, "building end")
    }
}

