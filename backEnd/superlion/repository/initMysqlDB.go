//@program: superlion
//@author: yanjl
//@create: 2023-09-07 20:11
package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var db *gorm.DB

func InitMysqlDB() error {
	var err error

	// dsn := "{}:piper_2021%wii@tcp({{MYSQL_HOST}}:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "jnfdroot:19aa5b459e809559@tcp(mysql.sqlpub.com:3306)/jnfdcome?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 打印sql
		Logger: logger.Default.LogMode(logger.Info),
	})

	sqlDB, err := db.DB()
	if err != nil {
		log.Print("connect db server failed.", err.Error())
		return err
	}
	sqlDB.SetMaxIdleConns(8)                    // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxOpenConns(8)                    // SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetConnMaxLifetime(time.Second * 600) // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	// db = db

	fmt.Println("mysql init over,everything is OK")
	return err

}

func getDBHandler() *gorm.DB {
	return db
}
