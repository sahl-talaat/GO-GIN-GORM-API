package Routes

import (
	"myapp/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	v1 := r.Group("api/v1")
	{
		v1.GET("employees", Controllers.GetAllEmployees)
		v1.POST("employees", Controllers.CreateEmployee)
		v1.GET("employees/:id", Controllers.GetEmployee)
		v1.PUT("employees/:id", Controllers.UpdateEmployee)
		v1.DELETE("employees/:id", Controllers.DeleteEmployee)

		v1.GET("departments", Controllers.GetAllDepartment)
		v1.POST("departments", Controllers.CreateDepartment)
		v1.GET("departments/:id", Controllers.GetDepartment)
		v1.PUT("departments/:id", Controllers.UpdatedDepartment)
		v1.DELETE("departments/:id", Controllers.DeleteDepartment)
	}

}
