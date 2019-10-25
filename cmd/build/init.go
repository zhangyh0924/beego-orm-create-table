package build

import (
    "fmt"
    "github.com/astaxie/beego/config"
    "github.com/astaxie/beego/orm"
    _"github.com/go-sql-driver/mysql"
    "os"
    "reflect"
    "strings"
)

type DatabaseConf struct {
    Host     string
    Port     string
    Database string
    UserName string
    Password string
    Charset  string
    Prefix   string
}

func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)
    databaseConf := NewDatabaseConf().GetMysqlConfigStr()
    orm.RegisterDataBase("default", "mysql", databaseConf)
}

func NewDatabaseConf() *DatabaseConf  {
    return &DatabaseConf{}
}

func (d *DatabaseConf)GetMysqlConfigStr() string {
    d.GetMysqlConf()
    mysqlstr := d.UserName + ":" + d.Password + "@tcp(" + d.Host + ":" + d.Port + ")/information_schema?charset=" + d.Charset
    return mysqlstr
}

func (d *DatabaseConf) GetMysqlConf() *DatabaseConf {
    configer, e := config.NewConfig("ini", "./conf/database.conf")
    if e != nil {
        fmt.Println(e)
        os.Exit(-1)
    }
    typeof := reflect.TypeOf(d).Elem()
    valueOf := reflect.ValueOf(d).Elem()

    for i := 0; i < typeof.NumField(); i++ {
        name := strings.ToLower(typeof.Field(i).Name)
        value := configer.String(name)
        valueOf.Field(i).SetString(value)
    }
    return d
}
