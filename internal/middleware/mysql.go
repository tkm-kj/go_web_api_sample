package middleware

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var mysqlInstance *gorm.DB

func GetMySQLConnection() *gorm.DB {
	return mysqlInstance
}

// MySQL接続解除はmain関数の中で行ってあげること
func CloseMySQLConnection() {
	mysqlInstance.Close()
}

func init() {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME")))
	if err != nil {
		log.Fatalf("fatal error!: %+v", err)
	}
	db.LogMode(true)

	mysqlInstance = db
}
