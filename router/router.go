package router

import (
	"fmt"
	"net/http"
	"scyllax-pjp/controller"
	"scyllax-pjp/exception"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(
	pjpController *controller.PjpController,
) *gin.Engine {
	service := gin.Default()

	service.Use(gin.CustomRecovery(func(c *gin.Context, anyError any) {
		err := WrapErr(anyError)
		// change with ur own error handle here
		exception.ErrorHandler(err, c, nil)
	}))

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

func WrapErr(err any) error {
	if err == nil {
		return nil
	}

	var errs error
	switch e := err.(type) {
	case error:
		errs = e
	default:
		errs = fmt.Errorf("%v", e)
	}

	return errs
}
