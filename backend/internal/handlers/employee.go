package handlers

import (
	"net/http"
	"project-backend/internal/database"
	"project-backend/internal/models"

	"github.com/gin-gonic/gin"
)

// GetEmployees retrieves all employees with optional department information
func GetEmployees(c *gin.Context) {
	var employees []models.Employee

	// Query with department preloading
	result := database.DB.Preload("Department").Find(&employees)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  employees,
		"count": len(employees),
	})
}

// GetEmployee retrieves a single employee by ID
func GetEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee

	result := database.DB.Preload("Department").First(&employee, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": employee})
}

// CreateEmployee creates a new employee
func CreateEmployee(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&employee)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Reload with department information
	database.DB.Preload("Department").First(&employee, employee.ID)

	c.JSON(http.StatusCreated, gin.H{"data": employee})
}

// UpdateEmployee updates an existing employee
func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee

	// Check if employee exists
	if err := database.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	// Bind updated data
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update employee
	if err := database.DB.Save(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Reload with department information
	database.DB.Preload("Department").First(&employee, employee.ID)

	c.JSON(http.StatusOK, gin.H{"data": employee})
}

// DeleteEmployee soft deletes an employee
func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee

	result := database.DB.First(&employee, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	if err := database.DB.Delete(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}

// GetEmployeesByDepartment retrieves employees by department ID
func GetEmployeesByDepartment(c *gin.Context) {
	departmentID := c.Param("departmentId")
	var employees []models.Employee

	result := database.DB.Preload("Department").Where("department_id = ?", departmentID).Find(&employees)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  employees,
		"count": len(employees),
	})
}

// GetEmployeesByStatus retrieves employees by status
func GetEmployeesByStatus(c *gin.Context) {
	status := c.Query("status")
	if status == "" {
		status = "active" // default to active
	}

	var employees []models.Employee
	result := database.DB.Preload("Department").Where("status = ?", status).Find(&employees)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   employees,
		"count":  len(employees),
		"status": status,
	})
}
