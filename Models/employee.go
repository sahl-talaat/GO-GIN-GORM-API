package models

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name         string
	Age          int
	Salary       float64
	DepartmentID uint       `gorm:"index"`
	Department   Department `gorm:"foreignkey:DepartmentID"`
}
