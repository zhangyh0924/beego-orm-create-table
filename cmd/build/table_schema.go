package build

import (
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
    "strings"
)

type TableSchema struct {
    TableName   string
    TableColumn string
}

func NewTableSchema() *TableSchema {
    return &TableSchema{}
}

func (this *TableSchema) GetAll(tableName string) []orm.Params {
    query := orm.NewOrm()
    sql := "SELECT TABLE_NAME,TABLE_COMMENT FROM `TABLES` WHERE TABLE_SCHEMA='"+NewDatabaseConf().GetMysqlConf().Database+"'"
    if tableName != "" {
        sql += " AND TABLE_NAME='" + tableName + "'"
    }
    var tableSchemas []orm.Params
    query.Raw(sql).Values(&tableSchemas)
    return tableSchemas
}

// 获取写入字符串
func (this *TableSchema) GetWriteStr(columns []*Columns) string {
    tool := NewTool()
    pack := "package models\n\n"
    isTime := false
    this.TableName = strings.TrimLeft(this.TableName, NewDatabaseConf().GetMysqlConf().Prefix)
    str := "type " + tool.FirstByteUp(this.TableName) + " struct {\n"
    long := Longest(columns)
    for _, val := range columns {
        structName := tool.FirstByteUp(val.ColumnName)
        blankSpace := getBlankSpace(long, len(structName))
        str = str + "    " + structName + blankSpace + tool.DataTypeSwitch(val.DataType) + "            `orm:\"column(" + val.ColumnName + ")"
        if val.ColumnKey == "PRI" {
            str = str + ";pk;auto;unique"
        }
        str = str + "\" json:\"" + val.ColumnName + "\"`    //" + val.ColumnComment + "\n"

        if tool.DataTypeSwitch(val.DataType) == "time.Time" {
            isTime = true
        }
    }
    if isTime {
        pack += "import (\n    \"time\"\n)\n\n"
    }
    str += "}"
    return pack + str + getNewFunc(tool.FirstByteUp(this.TableName))
}

// 最长字段长度
func Longest(columns []*Columns) int {
    lenth := 0
    for _, val := range columns {
        if l := len(val.ColumnName); l > lenth {
            lenth = l
        }
    }
    return lenth
}

//获取空格长度
func getBlankSpace(sourceLen, lenth int) string {
    str := "        "
    for i := 0; i < (sourceLen - lenth); i++ {
        str += " "
    }
    return str
}

func getNewFunc(name string) (str string) {
    str = ""
    str = "\n\nfunc New" + name + "() *" + name + " {\n"
    str += "    return &" + name + "{}\n}"
    return str
}
