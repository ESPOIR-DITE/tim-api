package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// const DSN string = "user.home.controller.domain.controller:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

//const dsn = "timtubeAdmin:1234@1234@tcp(127.0.0.1:3306)/timtubedb?charset=utf8mb4&parseTime=True&loc=Local"

// const dsn = "timtubeAdmin:123#Lunchbar@tcp(74.208.50.103:3306)/timtubedb?charset=utf8mb4&parseTime=True&loc=Local"
const dsn = "timtubeAdmin@localhost:123#Lunchbar@tcp(localhost:3306)/timtubedb?charset=utf8mb4&parseTime=True&loc=Local"

//const DOCKERDSN = "timtubeAdmin:1234@tcp(172.17.0.3:3306)/timtubedb?charset=utf8mb4&parseTime=True&loc=Local"

func GetDatabase() (db *gorm.DB) {
	//dsn := "tcp://localhost:9000?database=gorm&username=gorm&password=gorm&read_timeout=10&write_timeout=20"
	//db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//dsn := "timtubeAdmin:1234@1234@tcp(127.0.0.1:3306)/timtubedb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	sqlDB, err := db.DB()
	sqlDB.SetConnMaxLifetime(time.Minute)
	return db
}
