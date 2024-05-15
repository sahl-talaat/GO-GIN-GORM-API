package Models

type Employee struct {
	//gorm.Model // provides fields for 'ID', 'CreatedAt', 'UpdatedAt', and 'DeletedAt
	ID           uint `gorm:"primarykey"`
	Salary       float64
	Name         string
	Age          int
	DepartmentID uint `gorm:"index"` // Foreign key referencing Department's primary key
	// Define foreign key relationship
	Department Department `gorm:"foreignkey:DepartmentID"`
	// Department field is added to represent the relationship with the Department model.
	//gorm:"foreignkey:DepartmentID" specifies that
	//DepartmentID in the Employee struct is the foreign key that references the ID field of the Department struct.
}

type Department struct {
	ID   uint
	Name string
}
