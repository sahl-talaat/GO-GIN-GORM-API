package models

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model // provides fields for 'ID', 'CreatedAt', 'UpdatedAt', and 'DeletedAt
	//ID   int
	Salary       float64
	DepartmentID uint `gorm:"index"` // Foreign key referencing Department's primary key
	// Define foreign key relationship
	Department Department `gorm:"foreignkey:DepartmentID"`
	// Department field is added to represent the relationship with the Department model.
	//gorm:"foreignkey:DepartmentID" specifies that
	//DepartmentID in the Employee struct is the foreign key that references the ID field of the Department struct.
	PersonalDataID uint         `gorm:"index"`
	PersonalData   PersonalData `gorm:"foreignkey:PersonalDataID"`

	WorkSessionID uint        `gorm:"index"`
	WorkSession   WorkSession `gorm:"foreignkey:WorkSessionID"`
}

/* func (emp *Employee) SetStartTime() error {
	workSession := WorkSession{
		StartTime: time.Now(),
		Model:     emp.Model,
	}
	return database.GetDB().Create(&workSession).Error
}

func (emp *Employee) SetEndTime() error {
	// Get the latest work session for the employee
	var workSession WorkSession
	if err := database.GetDB().Last(&workSession, "employee_id = ?", emp.Model).Error; err != nil {
		return err
	}

	// Update the end time
	workSession.EndTime = time.Now()

	// Calculate hours worked
	workSession.CalculateHoursWorked()

	// Save updated work session
	return database.GetDB().Save(&workSession).Error
} */
