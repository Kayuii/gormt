package gtools

const (
	_tagGorm = "gorm"
	_tagJSON = "json"
)

//TypeDicMp 精确匹配类型
var TypeDicMp = map[string]string{
	"bit(1)":              "bool",
	"tinyint(1)":          "bool",
	"tinyint(1) unsigned": "bool",
	"tinyint(4)":          "int8",
	"json":                "string",

	"int":                "int",
	"integer":            "int",
	"tinyint":            "int",
	"smallint":           "int",
	"mediumint":          "int",
	"bigint":             "int64",
	"int unsigned":       "int",
	"integer unsigned":   "int",
	"tinyint unsigned":   "int",
	"smallint unsigned":  "int",
	"mediumint unsigned": "int",
	"bigint unsigned":    "int64",
	"bit":                "int",
	"bool":               "bool",
	"enum":               "string",
	"set":                "string",
	"varchar":            "string",
	"char":               "string",
	"tinytext":           "string",
	"mediumtext":         "string",
	"text":               "string",
	"longtext":           "string",
	"blob":               "string",
	"tinyblob":           "string",
	"mediumblob":         "string",
	"longblob":           "string",
	"date":               "time.Time",
	"datetime":           "time.Time",
	"timestamp":          "time.Time",
	"time":               "time.Time",
	"float":              "float64",
	"double":             "float64",
	"decimal":            "float64",
	"binary":             "string",
	"varbinary":          "string",
}

//TypeMatchMp 模糊匹配类型
var TypeMatchMp = map[string]string{
	`^(int)[(]\d+[)]`:         "int",
	`^(bigint)[(]\d+[)]`:      "int64",
	`^(char)[(]\d+[)]`:        "string",
	`^(varchar)[(]\d+[)]`:     "string",
	`^(float)[(]\d+,\d+[)]`:   "float64",
	`^(double)[(]\d+[)]`:      "float64",
	`^(decimal)[(]\d+,\d+[)]`: "float64",
}
