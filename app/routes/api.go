package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"log"
	"net/http"
	"testProject/app/controllers"
	"testProject/app/models"
	_ "testProject/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.

// @host localhost:8080
// @BasePath /

func SetupRouter(db *gorm.DB) {
	router := gin.Default()
	err := db.AutoMigrate(&models.Song{})
	if err != nil {
		log.Fatal(err)
	}

	// Swagger 2.0 Meta Information
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")
	{
		api.GET("/create", controllers.CreateSongController{DB: db}.CreateSong)
		api.GET("/update", controllers.UpdateSongController{DB: db}.UpdateSong)
		api.GET("/delete", controllers.DeleteSongController{DB: db}.DeleteSong)
		api.GET("/getPlayList", controllers.GetPlayListController{DB: db}.GetPlayList)
		api.GET("/getTrackText", controllers.GetTrackTextController{DB: db}.GetTrackText)
		api.GET("/info", controllers.GetDetailInfoSongController{DB: db}.GetDetailInfoSong)
		api.GET("/data", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "This is protected data"})
		})
	}

	err = router.Run()
	if err != nil {
		panic(err)
	}
}
