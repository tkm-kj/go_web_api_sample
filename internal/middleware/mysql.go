package middleware

import (
	"log"

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
	db, err := gorm.Open("mysql", "root:password@tcp(db:3306)/dev_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("fatal error!: %+v", err)
	}
	db.LogMode(true)

	mysqlInstance = db
}
