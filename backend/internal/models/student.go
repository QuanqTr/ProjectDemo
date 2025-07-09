package models

import (
	"time"
	"gorm.io/gorm"
)

type Student struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	StudentCode string         `json:"student_code" gorm:"unique;not null;size:20"`
	FirstName   string         `json:"first_name" gorm:"not null;size:100"`
	LastName    string         `json:"last_name" gorm:"not null;size:100"`
	Email       string         `json:"email" gorm:"unique;not null;size:255"`
	Phone       *string        `json:"phone" gorm:"size:20"`
	DateOfBirth *time.Time     `json:"date_of_birth"`
	Address     *string        `json:"address" gorm:"size:500"`
	Major       *string        `json:"major" gorm:"size:100"`
	Year        *int           `json:"year" gorm:"check:year >= 1 AND year <= 6"`
	GPA         *float64       `json:"gpa" gorm:"type:decimal(3,2);check:gpa >= 0 AND gpa <= 4"`
	Status      string         `json:"status" gorm:"default:'active';size:20;check:status IN ('active', 'inactive', 'graduated', 'suspended')"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// TableName specifies the table name for Student model
func (Student) TableName() string {
	return "students"
}
