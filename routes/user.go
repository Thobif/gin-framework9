package routes

import (
	"github.com/gin-gonic/gin"
	usercontroller "github.com/github.com/Thobif/gin-framework9/controllers/user"
)

func InitUserRoutes(rg *gin.RouterGroup) {
	routerGroup:= rg.Group("/users")

	//{domain_url}/api/v1/users
	routerGroup.GET("/",usercontroller.GetAll)
	//{domain_url}/api/v1/users/register
	routerGroup.POST("/register",usercontroller.Register)
	//{domain_url}/api/v1/users/login
	routerGroup.POST("/login",usercontroller.Login)
	//{domain_url}/api/v1/users/1
	routerGroup.GET("/:id",usercontroller.GetById)
	//{domain_url}/api/v1/users/search?fullname=Mark
	routerGroup.GET("/search",usercontroller.SearchByFullName)
}