package handlers

import (
	"net/http"
	"project-backend/internal/database"
	"project-backend/internal/models"

	"github.com/gin-gonic/gin"
)

// GetDepartments retrieves all departments with manager information
func GetDepartments(c *gin.Context) {
	var departments []models.Department
	result := database.DB.Preload("Manager").Find(&departments)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  departments,
		"count": len(departments),
	})
}

// GetDepartment retrieves a single department by ID with manager and employees
func GetDepartment(c *gin.Context) {
	id := c.Param("id")
	var department models.Department

	result := database.DB.Preload("Manager").Preload("Employees").First(&department, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Department not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": department})
}

// UpdateDepartment updates an existing department
func UpdateDepartment(c *gin.Context) {
	id := c.Param("id")
	var department models.Department

	// Check if department exists
	if err := database.DB.First(&department, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Department not found"})
		return
	}

	// Bind updated data
	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update department
	if err := database.DB.Save(&department).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Reload with manager information
	database.DB.Preload("Manager").First(&department, department.ID)

	c.JSON(http.StatusOK, gin.H{"data": department})
}

// CreateDepartment creates a new department
func CreateDepartment(c *gin.Context) {
	var department models.Department
	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&department)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": department})
}
