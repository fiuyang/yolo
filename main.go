package main

import (
	"scyllax-pjp/config"
	"scyllax-pjp/controller"
	_ "scyllax-pjp/docs"
	"scyllax-pjp/helper"
	"scyllax-pjp/migrations"
	"scyllax-pjp/repository"
	"scyllax-pjp/router"
	"scyllax-pjp/service"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

// @title 	Boilerplate API
// @version	1.0
// @description A Boilerplate API in Go using Gin framework

// @host 	localhost:8888
// @BasePath /api

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	//Database
	db := config.ConnectionDB(&loadConfig)

	//Validation
	validate := validator.New()
	_ = validate.RegisterValidation("unique", func(fl validator.FieldLevel) bool {
		return helper.ValidateUnique(db, fl)
	})

	//Migrations
	migrations.AutoMigrate(db)

	//Init Repository
	pjpRepository := repository.NewPjpRepositoryImpl(db)

	//Init Service
	pjpService := service.NewPjpServiceImpl(pjpRepository)

	//Init controller
	pjpController := controller.NewPjpController(pjpService, validate)

	//Router
	routes := router.NewRouter(
		pjpController,
	)

	server := &http.Server{
		Addr:           ":" + loadConfig.ServerPort,
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	error := server.ListenAndServe()
	helper.ErrorPanic(error)
}
