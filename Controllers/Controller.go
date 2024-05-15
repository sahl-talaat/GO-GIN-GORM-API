package Controllers

import (
	"myapp/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllEmployees(c *gin.Context) {
	var employee []Models.Employee
	if err := Models.GetAllEmployees(&employee); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	} else {
		c.JSON(http.StatusOK, employee)
	}
}

func CreateEmployee(c *gin.Context) {
	var employee Models.Employee
	if err := c.BindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, employee)
		return
	}
	if err := Models.CreateEmployee(&employee); err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, employee)
	}
}

func GetEmployee(c *gin.Context) {
	employeeID := c.Param("id")
	var employee Models.Employee
	if err := Models.GetEmployee(&employee, employeeID); err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, employee)
	}
}

func UpdateEmployee(c *gin.Context) {
	employeeID := c.Param("id")
	var employee Models.Employee
	if err := c.BindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, employee)
		return
	}
	if err := Models.UpdateEmployee(&employee, employeeID); err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, employee)
	}
}

func DeleteEmployee(c *gin.Context) {
	employeeID := c.Param("id")
	var employee Models.Employee
	if err := Models.DeleteEmployee(&employee, employeeID); err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"id " + employeeID: "deleted"})
	}
}

// department

func GetDepartment(c *gin.Context) {
	departmentID := c.Param("id")
	var deparetment Models.Department
	if err := Models.GetDepartment(&deparetment, departmentID); err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, deparetment)
	}
}

func CreateDepartment(c *gin.Context) {
	departmentID := c.Param("id")
	var department Models.Department
	if err := c.BindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, department)
		return
	}
	if err := Models.GetDepartment(&department, departmentID); err == nil { // equal equal
		c.JSON(http.StatusBadRequest, gin.H{"error": "department:" + departmentID + "existing alredy"})
		return
	}
	if err := Models.CreateDepartment(&department); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	} else {
		c.JSON(http.StatusOK, department)
	}
}

func GetAllDepartment(c *gin.Context) {
	var department []Models.Department
	if err := Models.GetAllDepartment(&department); err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, department)
	}
}

func UpdatedDepartment(c *gin.Context) {
	departmentID := c.Param("id")
	var department Models.Department
	if err := c.BindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, department)
		return
	}
	if err := Models.UpdatedDepartment(&department, departmentID); err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, department)
	}
}

func DeleteDepartment(c *gin.Context) {
	departmentID := c.Param("id")
	var department Models.Department
	if err := Models.DeleteDepartment(&department, departmentID); err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "  id: " + departmentID + " deleted"})
	}
}
