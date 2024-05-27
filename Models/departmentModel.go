package models

import (
	config "myapp/Config"
	"time"
)

func GetDepartments(department *[]Department) (err error) {
	err = config.DB.Preload("Employees").Find(&department).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateDepartment(department *Department) (err error) {
	err = config.DB.Create(&department).Error
	if err != nil {
		return err
	}
	return nil
}

func GetDepartment(department *Department, id string) (err error) {
	err = config.DB.Preload("Employees").Where("id = ?", id).First(&department).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateDepartment(department *Department, id string) (err error) {
	currentTime := time.Now()
	var isDepartmentExist Department
	err = GetDepartment(&isDepartmentExist, id)
	if err != nil {
		return err
	}
	isDepartmentExist.Name = department.Name
	isDepartmentExist.CreatedAt = currentTime
	err = config.DB.Save(&isDepartmentExist).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteDepartment(department *Department, id string) (err error) {
	err = config.DB.Preload("Employees").First(department, id).Error
	if err != nil {
		return err
	}
	if len(department.Employees) > 0 {
		err = config.DB.Model(&Employee{}).Where("department_id = ?", id).Update("department_id = ?", nil).Error
		if err != nil {
			return err
		}
	}
	err = config.DB.Delete(&Department{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
