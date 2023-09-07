package mysqlservice

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init() {
	// 初始化MySQL连接
	dsn := "root:Yan123456@tcp(yanjl.eu.org:9556)/superlion"
	d, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db = d
}

type User struct {
	ID   int
	Name string
}

func Mysql() {
	Init()
	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {

		// 查询用户
		rows, err := db.Query("SELECT id, name FROM users")
		fmt.Println("-----@@", rows)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// 解析和响应
		var users []User
		for rows.Next() {
			var u User
			err := rows.Scan(&u.ID, &u.Name)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			users = append(users, u)
		}
		c.JSON(200, gin.H{"users": users})
	})

	r.Run()
}

// func init() {
// 	panic("unimplemented")
// }
