//@program: superlion
//@author: yanjl
//@create: 2023-09-07 20:11
package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitMysqlDB() error {
	var err error

	//dsn := "{}:piper_2021%wii@tcp({{MYSQL_HOST}}:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "jnfdroot:19aa5b459e809559@tcp(mysql.sqlpub.com:3306)/jnfdcome?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return err

}
