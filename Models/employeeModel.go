package models

import (
	config "myapp/Config"
	"strconv"
)

func GetEmployees(employee *[]Employee) (err error) {
	if err = config.DB.Preload("Department").Find(employee).Error; err != nil {
		return err
	}
	return nil
}

func CreateEmployee(employee *Employee) (err error) {
	var isDepartmentExist Department
	departmentID := strconv.Itoa(int(employee.DepartmentID))
	if err = GetDepartment(&isDepartmentExist, departmentID); err != nil {
		return err
	}
	isDepartmentExist.Employees = append(isDepartmentExist.Employees, *employee)
	if err = config.DB.Create(employee).Error; err != nil {
		return err
	}
	employyID := strconv.Itoa(int(employee.ID))
	if err = GetEmployee(employee, employyID); err != nil {
		return err
	}
	return nil
}

func GetEmployee(employee *Employee, id string) (err error) {
	if err = config.DB.Preload("Department").Where("id = ?", id).First(&employee).Error; err != nil {
		return err
	}
	return nil
}

func UpdateEmployee(employee *Employee, id string) (err error) {
	var isEmployeeExist Employee
	if err = GetEmployee(&isEmployeeExist, id); err != nil {
		return err
	}
	departmentID := strconv.Itoa(int(employee.DepartmentID))
	var isDepartmentExist Department
	if err = GetDepartment(&isDepartmentExist, departmentID); err != nil {
		return err
	}
	isEmployeeExist.Name = employee.Name
	isEmployeeExist.Age = employee.Age
	isEmployeeExist.Salary = employee.Salary
	isEmployeeExist.DepartmentID = employee.DepartmentID
	isEmployeeExist.Department = isDepartmentExist
	if err = config.DB.Save(&isEmployeeExist).Error; err != nil {
		return err
	}
	if err = GetEmployee(employee, id); err != nil {
		return err
	}
	return nil
}

func DeleteEmployee(employee *Employee, id string) (err error) {
	if err = GetEmployee(employee, id); err != nil {
		return err
	}
	if err = config.DB.Where("id = ?", id).Delete(employee).Error; err != nil {
		return err
	}
	return nil
}
