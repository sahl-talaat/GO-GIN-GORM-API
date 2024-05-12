package main

import (
	database "example.com/myproject/Desktop/sahl/golang_learn/apps/emp_sys_gin_gorm/DataBase"
	//models "example.com/myproject/Desktop/sahl/golang_learn/apps/emp_sys_gin_gorm/Models"
)

func main() {

	_, err := database.InitializeDB()
	if err != nil {
		panic(err)
	}

	//db := database.GetDB()

	/* department := models.Department{}
	db.Create(&department)

	personalData := models.PersonalData{}
	db.Create(&personalData)

	workSession := models.WorkSession{}
	db.Create(&workSession)

	employee := models.Employee{DepartmentID: department.ID, PersonalDataID: personalData.ID, WorkSessionID: workSession.ID}
	db.Create(&employee) */

}
