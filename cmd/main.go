package main

import (
    "cmd/build"
    "flag"
    "fmt"
    "os"
)

var h bool
var buildFile *build.Build

func init() {
    buildFile = build.NewBuild()
    flag.BoolVar(&h, "h", false, "生成表的名称");
    flag.StringVar(&buildFile.TableName, "t", "", "生成表的名称");
    flag.StringVar(&buildFile.Path, "p", "", "文件保存路径");
    flag.BoolVar(&buildFile.IsFileHump, "m", false, "文件是否驼峰 true false 默认false");
    flag.Parse()
}

func main() {
    if h {
        fmt.Fprintf(os.Stderr, `usage  rum main.go -t=tablename -m=ishump -p=filepath
option:
`)
        flag.PrintDefaults()
        return
    }
    e := buildFile.Build()
    if e != nil {
        fmt.Println(e)
        os.Exit(-1)
    }
    fmt.Println("over")
}
