package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	rModule       = "module"
	rProject      = "project"
	rFunction     = "function"
	rType         = "type"
	rName         = "name"
	InternalError = http.StatusForbidden
)

//TODO:
// update project controller and model
// testing postman

func ApplyRoutes(r *gin.Engine) {

	projects := r.Group("/" + rProject)
	{
		projects.GET("/list", listProject)
		projects.GET("/get", findProjectById)
		projects.GET("/get/:"+rProject, findProjectByName)
		projects.POST("/add", addProject)
		projects.PATCH("/update", updateProject)
		projects.DELETE("/delete", deleteProject)
	}
	modules := r.Group("/" + rModule)
	{
		modules.GET("/list", listModules)
		modules.GET("/get", findModuleById)
		modules.GET("/get/:"+rModule, findModuleByName)
		modules.POST("/add", addModule)
		modules.PATCH("/update", updateModule)
		modules.DELETE("/delete", deleteModule)
	}
	functions := r.Group("/" + rFunction)
	{
		functions.GET("/list", listFunctions)
		functions.GET("/get", findFunctionById)
		functions.GET("/get/:"+rFunction, findFunctionByName)
		functions.POST("/add", addFunction)
		functions.PATCH("/update", updateFunction)
		functions.DELETE("/delete", deleteFunction)
	}
	//types := r.Group("/" + rType)
	//{
	//	types.GET("/list", listTypes)
	//	types.GET("/get", findTypeById)
	//	types.GET("/get/:"+rType, findTypeByName)
	//	types.POST("/add", addType)
	//	types.PATCH("/update", updateType)
	//	types.DELETE("/delete", deleteType)
	//}
}