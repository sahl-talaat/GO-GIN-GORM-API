package models

import "gorm.io/gorm"

type PersonalData struct {
	gorm.Model
	Name string
	Age  int
}
