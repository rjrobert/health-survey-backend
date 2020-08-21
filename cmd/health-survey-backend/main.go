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

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
		v1.Use(middlewares.APIKeyAuth())
		users := v1.Group("/users")
		{
			users.GET(":id", apis.GetUser)
		}

		surveys := v1.Group("/surveys")
		{
			surveys.GET(":id", apis.GetSurveys)
		}
	}

	config.Config.DB, config.Config.DBErr = gorm.Open("postgres", config.Config.DSN)
	if config.Config.DBErr != nil {
		panic(config.Config.DBErr)
	}

	setupDatabase()

	defer config.Config.DB.Close()

	log.Println("Successfully connected to database")

	r.Run(fmt.Sprintf(":%v", config.Config.ServerPort))
}

func setupDatabase() {
	// config.Config.DB.DropTableIfExists(&models.Survey{})
	// config.Config.DB.DropTableIfExists(&models.Page{})
	// config.Config.DB.DropTableIfExists(&models.Element{})
	// config.Config.DB.DropTableIfExists(&models.Choice{})
	// config.Config.DB.DropTableIfExists(&models.Column{})
	// config.Config.DB.DropTableIfExists(&models.Row{})

	// This is needed for generation of schema for postgres image.
	config.Config.DB.AutoMigrate(&models.User{})
	config.Config.DB.AutoMigrate(&models.Survey{})
	config.Config.DB.AutoMigrate(&models.Page{})
	config.Config.DB.AutoMigrate(&models.Element{})
	config.Config.DB.AutoMigrate(&models.Choice{})
	config.Config.DB.AutoMigrate(&models.Column{})
	config.Config.DB.AutoMigrate(&models.Row{})

	// Create Survey data in db if it doesn't already exist
	// daos.PreloadSurveyData()
}
