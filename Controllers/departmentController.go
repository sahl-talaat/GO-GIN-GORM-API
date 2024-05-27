package controllers

import (
	models "myapp/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDepartments(c *gin.Context) {
	var department []models.Department
	if err := models.GetDepartments(&department); err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, department)
	}
}

func CreateDepartment(c *gin.Context) {
	departmentID := c.Param("id")
	var department models.Department
	if err := c.BindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, department)
		return
	}
	if err := models.GetDepartment(&department, departmentID); err == nil { // equal equal
		c.JSON(http.StatusBadRequest, gin.H{"error": "department:" + departmentID + "existing alredy"})
		return
	}
	if err := models.CreateDepartment(&department); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	} else {
		c.JSON(http.StatusOK, department)
	}
}

func GetDepartment(c *gin.Context) {
	departmentID := c.Param("id")
	var deparetment models.Department
	if err := models.GetDepartment(&deparetment, departmentID); err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, deparetment)
	}
}

func UpdateDepartment(c *gin.Context) {
	departmentID := c.Param("id")
	var department models.Department
	if err := c.BindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, department)
		return
	}
	if err := models.UpdateDepartment(&department, departmentID); err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, department)
	}
}

func DeleteDepartment(c *gin.Context) {
	departmentID := c.Param("id")
	var department models.Department
	if err := models.DeleteDepartment(&department, departmentID); err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "  id: " + departmentID + " deleted"})
	}
}
