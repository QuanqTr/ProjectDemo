package main

import (
	"log"
	"project-backend/internal/database"
	"project-backend/internal/models"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Connect to database
	database.Connect()

	// Auto migrate models
	database.DB.AutoMigrate(&models.Student{})

	// Seed students
	students := []models.Student{
		{
			StudentCode: "SV001",
			FirstName:   "Nguyễn",
			LastName:    "Văn An",
			Email:       "nguyen.van.an@student.edu.vn",
			Phone:       stringPtr("0901234567"),
			DateOfBirth: timePtr(time.Date(2002, 5, 15, 0, 0, 0, 0, time.UTC)),
			Address:     stringPtr("123 Đường ABC, Quận 1, TP.HCM"),
			Major:       stringPtr("Computer Science"),
			Year:        intPtr(3),
			GPA:         float64Ptr(3.75),
			Status:      "active",
		},
		{
			StudentCode: "SV002",
			FirstName:   "Trần",
			LastName:    "Thị Bình",
			Email:       "tran.thi.binh@student.edu.vn",
			Phone:       stringPtr("0902345678"),
			DateOfBirth: timePtr(time.Date(2003, 8, 22, 0, 0, 0, 0, time.UTC)),
			Address:     stringPtr("456 Đường XYZ, Quận 2, TP.HCM"),
			Major:       stringPtr("Information Technology"),
			Year:        intPtr(2),
			GPA:         float64Ptr(3.85),
			Status:      "active",
		},
		{
			StudentCode: "SV003",
			FirstName:   "Lê",
			LastName:    "Minh Cường",
			Email:       "le.minh.cuong@student.edu.vn",
			Phone:       stringPtr("0903456789"),
			DateOfBirth: timePtr(time.Date(2001, 12, 10, 0, 0, 0, 0, time.UTC)),
			Address:     stringPtr("789 Đường DEF, Quận 3, TP.HCM"),
			Major:       stringPtr("Software Engineering"),
			Year:        intPtr(4),
			GPA:         float64Ptr(3.92),
			Status:      "active",
		},
		{
			StudentCode: "SV004",
			FirstName:   "Phạm",
			LastName:    "Thu Dung",
			Email:       "pham.thu.dung@student.edu.vn",
			Phone:       stringPtr("0904567890"),
			DateOfBirth: timePtr(time.Date(2002, 3, 8, 0, 0, 0, 0, time.UTC)),
			Address:     stringPtr("321 Đường GHI, Quận 4, TP.HCM"),
			Major:       stringPtr("Data Science"),
			Year:        intPtr(3),
			GPA:         float64Ptr(3.68),
			Status:      "active",
		},
		{
			StudentCode: "SV005",
			FirstName:   "Hoàng",
			LastName:    "Văn Em",
			Email:       "hoang.van.em@student.edu.vn",
			Phone:       stringPtr("0905678901"),
			DateOfBirth: timePtr(time.Date(2003, 11, 25, 0, 0, 0, 0, time.UTC)),
			Address:     stringPtr("654 Đường JKL, Quận 5, TP.HCM"),
			Major:       stringPtr("Computer Science"),
			Year:        intPtr(2),
			GPA:         float64Ptr(3.45),
			Status:      "active",
		},
		{
			StudentCode: "SV006",
			FirstName:   "Vũ",
			LastName:    "Thị Phương",
			Email:       "vu.thi.phuong@student.edu.vn",
			Phone:       stringPtr("0906789012"),
			DateOfBirth: timePtr(time.Date(2000, 7, 18, 0, 0, 0, 0, time.UTC)),
			Address:     stringPtr("987 Đường MNO, Quận 6, TP.HCM"),
			Major:       stringPtr("Information Technology"),
			Year:        intPtr(4),
			GPA:         float64Ptr(3.88),
			Status:      "graduated",
		},
		{
			StudentCode: "SV007",
			FirstName:   "Đặng",
			LastName:    "Minh Quang",
			Email:       "dang.minh.quang@student.edu.vn",
			Phone:       stringPtr("0907890123"),
			DateOfBirth: timePtr(time.Date(2002, 9, 5, 0, 0, 0, 0, time.UTC)),
			Address:     stringPtr("147 Đường PQR, Quận 7, TP.HCM"),
			Major:       stringPtr("Artificial Intelligence"),
			Year:        intPtr(3),
			GPA:         float64Ptr(3.95),
			Status:      "active",
		},
		{
			StudentCode: "SV008",
			FirstName:   "Bùi",
			LastName:    "Thị Hương",
			Email:       "bui.thi.huong@student.edu.vn",
			Phone:       stringPtr("0908901234"),
			DateOfBirth: timePtr(time.Date(2003, 4, 12, 0, 0, 0, 0, time.UTC)),
			Address:     stringPtr("258 Đường STU, Quận 8, TP.HCM"),
			Major:       stringPtr("Cybersecurity"),
			Year:        intPtr(2),
			GPA:         float64Ptr(3.72),
			Status:      "active",
		},
	}

	for _, student := range students {
		var existingStudent models.Student
		if err := database.DB.Where("student_code = ?", student.StudentCode).First(&existingStudent).Error; err != nil {
			// Student doesn't exist, create it
			if err := database.DB.Create(&student).Error; err != nil {
				log.Printf("Error creating student %s %s: %v", student.FirstName, student.LastName, err)
			} else {
				log.Printf("Created student: %s %s (%s)", student.FirstName, student.LastName, student.StudentCode)
			}
		} else {
			log.Printf("Student %s already exists", student.StudentCode)
		}
	}

	log.Println("Student seeding completed!")
}

// Helper functions are now in utils.go
