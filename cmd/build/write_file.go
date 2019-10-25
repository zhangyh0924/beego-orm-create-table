package build

import (
    "bufio"
    "errors"
    "fmt"
    "os"
)

type WriteFile struct {
    Content string
    Path    string
}

func NewWriteFile() *WriteFile {
    return &WriteFile{}
}

func (this *WriteFile) WriteFile() error {
    file, e := os.OpenFile(this.Path, os.O_CREATE|os.O_RDWR, 0666)
    if e != nil {
        return e
    }
    writer := bufio.NewWriter(file)
    if writer == nil {
        return errors.New("文件写入失败")
    }
    _, e = writer.WriteString(this.Content)
    if e != nil {
        return e
    }
    writer.Flush()
    fmt.Println(this.Path, "is success")
    return nil
}

func (this *WriteFile) CheckDir() {
    info, e := os.Stat(this.Path)
    if e != nil {
        os.MkdirAll(this.Path, 0666)
        return
    }
    if !info.IsDir() {
        os.MkdirAll(this.Path, 0666)
    }
}
