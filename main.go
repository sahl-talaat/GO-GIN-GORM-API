package main

import (
	"fmt"

	"myapp/Config"
	"myapp/Routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	models "myapp/Models"
)

var err error

func main() {
	//dsn := "root:sahl1049@tcp(127.0.0.1:3306)/EmpDB?charset=utf8mb4&parseTime=True&loc=Local"
	//database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if Config.DB, err = gorm.Open(mysql.Open(Config.DbURL(Config.BuildDBConfig())), &gorm.Config{}); err != nil {
		fmt.Println("status: ", err)
	}

	Config.DB.AutoMigrate(&models.Employee{}, &models.Department{})

	app := gin.Default()

	Routes.SetupRouter(app)

	app.Run(":8088")

	// _, err := database.InitializeDB()
	// if err != nil {
	// 	panic(err.Error())
	// }

	// app := gin.Default()

	// app.POST("/api/employees", routing.CreateEmployee)
	// app.POST("/api/departments", routing.CreateDepartment)

	// app.GET("api/employees/:id", routing.GetEmpByID)
	// app.GET("api/departments/:id", routing.GetDepartByID)

	// app.GET("/api/departments", routing.GetAllDepart)
	// app.GET("/api/employees", routing.GetAllEmplyee)

	// app.PUT("/api/employees/:id", routing.UpdateEmpByID)
	// app.PUT("/api/departments/:id", routing.UpdateDepartByID)

	// app.DELETE("/api/employees/:id", routing.DeleteEmpByID)
	// app.DELETE("/api/departments/:id", routing.DeleteDepartByID)

	// app.Run(":8088")

}
