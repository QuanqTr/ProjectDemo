package handlers

import (
	"net/http"
	"project-backend/internal/database"
	"project-backend/internal/models"

	"github.com/gin-gonic/gin"
)

// GetStudents retrieves all students
func GetStudents(c *gin.Context) {
	var students []models.Student
	
	result := database.DB.Find(&students)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"data":  students,
		"count": len(students),
	})
}

// GetStudent retrieves a single student by ID
func GetStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	
	result := database.DB.First(&student, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": student})
}

// CreateStudent creates a new student
func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&student)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": student})
}

// UpdateStudent updates an existing student
func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	
	// Check if student exists
	if err := database.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	// Bind updated data
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update student
	if err := database.DB.Save(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": student})
}

// DeleteStudent soft deletes a student
func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	
	result := database.DB.First(&student, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	if err := database.DB.Delete(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}

// GetStudentsByMajor retrieves students by major
func GetStudentsByMajor(c *gin.Context) {
	major := c.Query("major")
	if major == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Major parameter is required"})
		return
	}
	
	var students []models.Student
	result := database.DB.Where("major = ?", major).Find(&students)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"data":  students,
		"count": len(students),
		"major": major,
	})
}

// GetStudentsByStatus retrieves students by status
func GetStudentsByStatus(c *gin.Context) {
	status := c.Query("status")
	if status == "" {
		status = "active" // default to active
	}
	
	var students []models.Student
	result := database.DB.Where("status = ?", status).Find(&students)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"data":   students,
		"count":  len(students),
		"status": status,
	})
}
