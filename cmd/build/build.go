package build

import (
    "fmt"
    "strings"
)

type Build struct {
    TableName  string
    Path       string
    IsFileHump bool
}

func NewBuild() *Build {
    return &Build{}
}

func (this *Build) Build() error {
    tableSchema := NewTableSchema()
    columns := NewColumns()
    writeFile := NewWriteFile()

    all := tableSchema.GetAll(this.TableName)
    writeFile.Path = this.GetPath()
    writeFile.CheckDir()
    for _, values := range all {
        columns.TableName = values["TABLE_NAME"].(string)
        info := columns.GetInfo() // 获取表相关信息
        tableSchema.TableName = columns.TableName
        str := tableSchema.GetWriteStr(info) // 获取写入字符串
        writeFile.Path = this.Path + this.GetFileName(values["TABLE_NAME"].(string))
        writeFile.Content = str
        err := writeFile.WriteFile()
        if err != nil {
            fmt.Println(err)
        }
    }
    return nil
}

//获取文件夹路径，如果不存在创建
func (this *Build) GetPath() string {
    if this.Path == "" {
        this.Path = "./model/"
    } else {
        strings.TrimRight(this.Path, "/")
        this.Path += "/"
    }
    return this.Path
}

func (this *Build) GetFileName(name string) string {
    name = strings.TrimLeft(name, NewDatabaseConf().GetMysqlConf().Prefix)
    if this.IsFileHump {
        name = NewTool().FirstByteUp(name)
    }
    name += ".go"
    return name
}
