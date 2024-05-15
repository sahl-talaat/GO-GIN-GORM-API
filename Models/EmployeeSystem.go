package Models

import (
	"myapp/Config"
	"strconv"
)

func GetAllEmployees(employee *[]Employee) (err error) {
	if err = Config.DB.Preload("Department").Find(employee).Error; err != nil {
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
	if err = Config.DB.Create(employee).Error; err != nil {
		return err
	}
	employyID := strconv.Itoa(int(employee.ID))
	if err = GetEmployee(employee, employyID); err != nil {
		return err
	}
	return nil
}

func GetEmployee(employee *Employee, id string) (err error) {
	if err = Config.DB.Preload("Department").Where("id = ?", id).First(&employee).Error; err != nil {
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
	if err = Config.DB.Save(&isEmployeeExist).Error; err != nil {
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
	if err = Config.DB.Where("id = ?", id).Delete(employee).Error; err != nil {
		return err
	}
	return nil
}

// Department
func CreateDepartment(department *Department) (err error) {
	if err := Config.DB.Create(&department).Error; err != nil {
		return err
	}
	return nil
}

func GetAllDepartment(department *[]Department) (err error) {
	if err := Config.DB.Find(&department).Error; err != nil {
		return err
	}
	return nil
}

func GetDepartment(department *Department, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(&department).Error; err != nil {
		return err
	}
	return nil
}

func UpdatedDepartment(department *Department, id string) (err error) {
	var isDepartmentExist Department
	if err := GetDepartment(&isDepartmentExist, id); err != nil {
		return err
	}
	isDepartmentExist.Name = department.Name
	if err := Config.DB.Save(&isDepartmentExist).Error; err != nil {
		return err
	}
	return nil
}

func DeleteDepartment(department *Department, id string) (err error) {
	if err := GetDepartment(department, id); err != nil {
		return err
	}
	var employee []Employee
	if err := Config.DB.Preload("Department").Where("department_id = ?", id).Find(&employee).Error; err != nil {
		return err
	}
	if len(employee) > 0 {
		if err := Config.DB.Model(&Employee{}).Where("department_id = ?", id).Update("department_id", nil).Error; err != nil {
			return err
		}
	}
	if err := Config.DB.Delete(&department).Error; err != nil {
		return err
	}
	return nil
}
