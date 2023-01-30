package controllers

import "github.com/gin-gonic/gin"

func (pagecontroller *PageController) Routes(route *gin.RouterGroup) {
	route.GET("/pages", pagecontroller.GetAllPage)
	route.GET("/:query", pagecontroller.GetByQuery)
	route.POST("/newpage", pagecontroller.CreateNewPage)
	route.GET("/", online)

}
