package main

import (
	"log"
	"net/http"

	config "api-gateway/config"
	app_config "api-gateway/config/app_config"
	minioconfig "api-gateway/config/minio_config"
	services_config "api-gateway/config/services_config"
	"api-gateway/controllers"
	ApiKey "api-gateway/middlewares"
	"api-gateway/response"
	"api-gateway/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load konfigurasi service dari YAML
	cfg, err := services_config.LoadConfigServices("./config_services.yaml")
	if err != nil {
		log.Fatalf("Failed to load service config: %v", err)
	}

	app := gin.Default()
	config.InitConfig()

	// Init Minio client
	minioClient := minioconfig.NewMinioClient()
	fileService := utils.NewFileService(minioClient.Client, minioClient.Bucket)
	fileController := controllers.NewFileHandler(fileService)
	app.GET("/download", fileController.Download)
	app.POST("/upload", fileController.Upload)
	app.DELETE("/delete", fileController.Delete)

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.SuccessOneResponse(nil, "OK", http.StatusOK))
	})

	// Apply API KEY middleware globally
	app.Use(ApiKey.APIKeyMiddleware())
	app.Use(services_config.ProxyHandler(cfg))

	log.Default().Println("Server is running on ", app_config.AppEnv)
	if err := app.Run(app_config.AppPort); err != nil {
		log.Fatalf("failed to start app: %v", err)
	}
}
