// @program: superlion
// @author: yanjl
// @create: 2023-09-07 20:11
package repository

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/go-sql-driver/mysql" // 导入 mysql 驱动
	gormmysql "gorm.io/driver/mysql" // 导入 mysql 驱动
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var db *gorm.DB

func InitMysqlDB() error {
	var err error

	// host = "yanjl.eu.org"; port = 9556
	// dsn := "{}:piper_2021%wii@tcp({{MYSQL_HOST}}:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "jnfdroot:19aa5b459e809559@tcp(mysql.sqlpub.com:3306)/jnfdcome?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "mysql://avnadmin:AVNS_t-uzz7L3n-xJ5h3xdHX@common-yanjl002.d.aivencloud.com:24583/lion?ssl-mode=REQUIRED"

	// 加载 CA 证书
	rootCertPool := x509.NewCertPool()
	dir, _ := os.Getwd()
	log.Printf("Loading CA certificate from file: %s", dir)
	pem, err := os.ReadFile("repository/ca-cert.pem")
	if err != nil {
		log.Fatalf("Failed to read CA certificate: %v", err)
	}

	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		log.Fatalf("Failed to append CA certificate")
	}

	// 创建 TLS 配置
	tlsConfig := &tls.Config{
		RootCAs: rootCertPool,
	}

	// 注册 TLS 配置
	err = mysql.RegisterTLSConfig("aiven", tlsConfig)
	if err != nil {
		log.Fatalf("Failed to register TLS config: %v", err)
	}

	dsn := "avnadmin:AVNS_t-uzz7L3n-xJ5h3xdHX@tcp(common-yanjl002.d.aivencloud.com:24583)/lion?tls=aiven"
	// mysql.Open(dsn).Name()
	gconfig := &gorm.Config{Logger: logger.Default.LogMode(logger.Info)}
	db, err = gorm.Open(gormmysql.Open(dsn), gconfig)
	if err != nil {
		log.Print("connect db server failed on step 1.", err.Error())
		return err
	}
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

/*
*
分页封装
*/
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
func getDBHandler() *gorm.DB {
	return db
}
