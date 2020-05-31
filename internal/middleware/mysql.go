package middleware

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tkm-kj/go_web_api_sample/internal/config"
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
	var db *gorm.DB
	var err error
	env := config.GetEnvVar()
	if config.IsLocal() {
		db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", env.DBUser, env.DBPassword, env.DBHost, env.DBName))
	} else {
		db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s", env.DBUser, env.DBPassword, env.DBHost, env.DBName))
	}
	if err != nil {
		log.Fatalf("fatal error!: %+v", err)
	}
	db.LogMode(true)

	mysqlInstance = db
}
