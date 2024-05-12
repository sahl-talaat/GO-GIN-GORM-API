package database

import (
	"fmt"

	models "example.com/myproject/Desktop/sahl/golang_learn/apps/emp_sys_gin_gorm/Models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitializeDB() (*gorm.DB, error) {
	dsn := "root:sahl1049@tcp(127.0.0.1:3306)/DB1?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		//panic("faild to connect DB")
		return nil, fmt.Errorf("faild to connect database: %v", err)
	}

	db = database

	AutoMigrate()

	return db, nil
}

func GetDB() *gorm.DB {
	return db
}

// define the schema
func AutoMigrate() {
	// Auto migrate table
	db.AutoMigrate(&models.Employee{})
	db.AutoMigrate(&models.Department{})
	db.AutoMigrate(&models.PersonalData{})
	db.AutoMigrate(&models.WorkSession{})
}
