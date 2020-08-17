package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/rjrobert/health-survey-backend/cmd/health-survey-backend/apis"
	"github.com/rjrobert/health-survey-backend/cmd/health-survey-backend/config"
	_ "github.com/rjrobert/health-survey-backend/cmd/health-survey-backend/docs"
	"github.com/rjrobert/health-survey-backend/cmd/health-survey-backend/middlewares"
	"github.com/rjrobert/health-survey-backend/cmd/health-survey-backend/models"
)

// @title Health Survey Backend Swagger API
// @version 1.0
// @description Swagger API for Health Survey Backend.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email jhrobert@pm.me

// @license.name MIT
// @license.url https://github.com/rjrobert/health-survey-backend/blob/master/LICENSE

// @BasePath /api/v1
func main() {
	// load application configurations
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	// Creates a router without any middleware by default
	r := gin.New()

	// Gloabl middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		v1.Use(middlewares.CheckJWT())
		v1.GET("/users/:id", apis.GetUser)
	}

	config.Config.DB, config.Config.DBErr = gorm.Open("postgres", config.Config.DSN)
	if config.Config.DBErr != nil {
		panic(config.Config.DBErr)
	}

	config.Config.DB.AutoMigrate(&models.User{}) // This is needed for generation of schema for postgres image.

	defer config.Config.DB.Close()

	log.Println("Successfully connected to database")

	r.Run(fmt.Sprintf(":%v", config.Config.ServerPort))
}
