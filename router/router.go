package router

import (
	"scyllax-pjp/controller"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(
	pjpController *controller.PjpController,
) *gin.Engine {
	service := gin.Default()

	//add swagger docs
	service.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": 404, "message": "Page not found"})
	})

	router := service.Group("/api")

	pjpRouter := router.Group("/pjp")
	pjpRouter.GET("", pjpController.FindAll)
	pjpRouter.GET("/:pjpId", pjpController.FindById)
	pjpRouter.POST("", pjpController.Create)
	pjpRouter.PATCH("/:pjpId", pjpController.Update)
	pjpRouter.DELETE("/:pjpId", pjpController.Delete)

	return service
}
