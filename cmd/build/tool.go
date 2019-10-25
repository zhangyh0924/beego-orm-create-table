package build

import "strings"

type Tool struct {
}

func NewTool() *Tool {
    return &Tool{}
}

func (Tool) FirstByteUp(str string) (result string) {
    if len(str) == 0 {
        return str
    }
    str = strings.ToLower(str)
    splits := strings.Split(str, "_")
    for _, val := range splits {
        result = result + strings.ToUpper(val[0:1]) + val[1:]
    }
    return result
}

func (Tool)DataTypeSwitch(str string) (result string) {
    source := map[string]string{
        "int":        "int32",
        "bigint":     "int64 ",
        "bool":       "bool",
        "varchar ":   "string",
        "char":       "string",
        "text":       "string",
        "mediumtext": "string",
        "longtext":   "string",
        "date":       "time.Time",
        "datetime":   "time.Time",
        "tinyint":    "int8",
        "smallint":   "int16",
        "enum":       "int16",
        "decimal":    "float32",
        "float":      "float32",
        "timestamp":  "time.Time",
    }
    if _, ok := source[str]; !ok {
        return "string"
    }

    return source[str]
}
