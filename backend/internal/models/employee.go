package models

import (
	"time"

	"gorm.io/gorm"
)

type EmployeeStatus string

const (
	EmployeeStatusActive    EmployeeStatus = "active"
	EmployeeStatusInactive  EmployeeStatus = "inactive"
	EmployeeStatusSuspended EmployeeStatus = "suspended"
)

type Employee struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	FirstName      string         `json:"first_name" gorm:"column:first_name;not null"`
	LastName       string         `json:"last_name" gorm:"column:last_name;not null"`
	Email          string         `json:"email" gorm:"unique;not null"`
	Phone          *string        `json:"phone" gorm:"column:phone"`
	DepartmentID   *uint          `json:"department_id" gorm:"column:department_id"`
	Position       *string        `json:"position" gorm:"column:position"`
	Status         EmployeeStatus `json:"status" gorm:"type:employee_status;default:active;not null"`
	FaceDescriptor any            `json:"face_descriptor,omitempty" gorm:"column:face_descriptor;type:vector"`
	JoinDate       time.Time      `json:"join_date" gorm:"column:join_date;default:now();not null"`
	EmployeeID     *string        `json:"employee_id" gorm:"column:employee_id;size:20"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// Relationships
	Department *Department `json:"department,omitempty" gorm:"foreignKey:DepartmentID"`
}

type Department struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"unique;not null"`
	Description *string        `json:"description"`
	ManagerID   *uint          `json:"manager_id" gorm:"column:manager_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// Relationships
	Manager   *Employee  `json:"manager,omitempty" gorm:"foreignKey:ManagerID"`
	Employees []Employee `json:"employees,omitempty" gorm:"foreignKey:DepartmentID"`
}

// TableName specifies the table name for Employee model
func (Employee) TableName() string {
	return "employees"
}

// TableName specifies the table name for Department model
func (Department) TableName() string {
	return "departments"
}
