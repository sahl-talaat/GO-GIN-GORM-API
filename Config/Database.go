package Config

import (
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

// DB Config represent db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbCinfig := DBConfig{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		DBName:   "EmpDB",
		Password: "sahl1049",
	}
	return &dbCinfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

/*
func InitializeDB() (*gorm.DB, error) {
	dsn := "root:sahl1049@tcp(127.0.0.1:3306)/EmpDB?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		//panic("faild to connect DB")
		return nil, fmt.Errorf("faild to connect database: %v", err)
	}

	//database.Logger.LogMode(1)
	DB = database

	AutoMigrate()

	return DB, nil
}

func GetDB() *gorm.DB {
	return DB
}

// define the schema
func AutoMigrate() {
	// Auto migrate table
	DB.AutoMigrate(&models.Employee{})
	DB.AutoMigrate(&models.Department{})
} */
