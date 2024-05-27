package controllers

import (
	models "myapp/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEmployees(c *gin.Context) {
	var employee []models.Employee
	if err := models.GetEmployees(&employee); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	} else {
		c.JSON(http.StatusOK, employee)
	}
}

func CreateEmployee(c *gin.Context) {
	var employee models.Employee
	if err := c.BindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, employee)
		return
	}
	if err := models.CreateEmployee(&employee); err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, employee)
	}
}

func GetEmployee(c *gin.Context) {
	employeeID := c.Param("id")
	var employee models.Employee
	if err := models.GetEmployee(&employee, employeeID); err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, employee)
	}
}

func UpdateEmployee(c *gin.Context) {
	employeeID := c.Param("id")
	var employee models.Employee
	if err := c.BindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, employee)
		return
	}
	if err := models.UpdateEmployee(&employee, employeeID); err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, employee)
	}
}

func DeleteEmployee(c *gin.Context) {
	employeeID := c.Param("id")
	var employee models.Employee
	if err := models.DeleteEmployee(&employee, employeeID); err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"id " + employeeID: "deleted"})
	}
}
