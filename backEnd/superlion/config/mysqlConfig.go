//@program: superlion
//@author: yanjl
//@create: 2023-09-26 14:38
package config

import (
	"encoding/json"
	"os"
)

// Config 配置对象
type Config struct {
	Database *Database `json:"database"`
}

// GlobalConfigSetting 配置实例
var GlobalConfigSetting = &Config{}

// Setup 配置
func Setup() error {
	filePtr, err := os.Open("config/config.json") //config的文件目录
	if err != nil {
		return err
	}
	defer filePtr.Close()
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(GlobalConfigSetting)
	DatabaseSetting = GlobalConfigSetting.Database

	return nil
}

// Database 数据库配置对象
type Database struct {
	Type        string `json:"type"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Host        string `json:"host"`
	Port        string `json:"port"`
	Name        string `json:"name"`
	TablePrefix string `json:"table_prefix"`
}

// DatabaseSetting 数据库配置对象 实例
var DatabaseSetting = &Database{}
