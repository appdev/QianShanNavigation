package router

import (
	"github.com/gin-gonic/gin"
	"goNav/controller"
	"goNav/middleware"
)

func InitRouter(router *gin.Engine) *gin.Engine {
	//router.GET("/users", Users)
	//router.PUT("/user/:id", Update)
	//router.DELETE("/user/:id", Destroy)
	router.POST("/login", controller.Login)
	router.POST("/addweb", middleware.JWTAuth(), controller.AddWeb)
	router.POST("/add/category", middleware.JWTAuth(), controller.AddCategory)
	router.GET("/delete", middleware.JWTAuth(), controller.Delete)
	router.POST("/update", middleware.JWTAuth(), controller.Update)
	router.POST("/update/category", middleware.JWTAuth(), controller.UpdateCategory)
	router.GET("/finds", middleware.JWTAuth(), controller.Finds)
	router.GET("/background", controller.GetNewBg)
	router.GET("/background/change", middleware.JWTAuth(), controller.GetBg)
	router.GET("/favicon", middleware.JWTAuth(), controller.Favicon)
	return router
}
