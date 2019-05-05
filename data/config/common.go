package config

import (
	"fmt"
	"public/tools"

	"github.com/BurntSushi/toml"
)

//
type CfgBase struct {
	SerialNumber       string `json:"serial_number" toml:"serial_number"`             //版本号
	ServiceName        string `json:"service_name" toml:"service_name"`               //service名字
	ServiceDisplayname string `json:"service_displayname" toml:"service_displayname"` //显示名
	SerciceDesc        string `json:"sercice_desc" toml:"sercice_desc"`               //service描述
	IsDev              bool   `json:"is_dev" toml:"is_dev"`                           //是否是开发版本
}

var _map = Config{}

func init() {
	onInit()
}

func onInit() {
	path := tools.GetModelPath()
	err := InitFile(path + "/config.toml")
	if err != nil {
		fmt.Println("InitFile: ", err.Error())
		return
	}
}

//OnIsDev ... 是否是开发版本
func OnIsDev() bool {
	return _map.IsDev
}

//InitFile ...
func InitFile(filename string) error {
	if _, err := toml.DecodeFile(filename, &_map); err != nil {
		fmt.Println("read toml error: ", err.Error())
		return err
	}

	return nil
}

//获取service配置信息
func GetServiceConfig() (name, displayName, desc string) {
	name = _map.ServiceName
	displayName = _map.ServiceDisplayname
	desc = _map.SerciceDesc
	return
}
