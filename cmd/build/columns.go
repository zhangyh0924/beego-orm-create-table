package build

import (
    "github.com/astaxie/beego/orm"
    "reflect"
)


type Columns struct {
    TableName     string
    ColumnName    string
    ColumnKey     string
    DataType      string
    Extra         string
    ColumnComment string
}

func NewColumns() *Columns {
    return &Columns{}
}

func (this *Columns) GetInfo() []*Columns {
    query := orm.NewOrm()
    sql := "SELECT COLUMN_NAME, DATA_TYPE, EXTRA,COLUMN_COMMENT, COLUMN_KEY FROM `COLUMNS` WHERE TABLE_NAME='" + this.TableName + "' AND TABLE_SCHEMA='card' ORDER BY ORDINAL_POSITION ASC"
    var values []orm.Params
    query.Raw(sql).Values(&values)
    var columns []*Columns
    for _, v := range values {
        column := &Columns{}
        valueOf := reflect.ValueOf(column).Elem()
        for key, val := range v {
            key = NewTool().FirstByteUp(key)
            valueOf.FieldByName(key).SetString(val.(string))
        }
        columns = append(columns, column)
    }

    return columns
}

