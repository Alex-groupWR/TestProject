package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"testProject/app/models"
	"testProject/app/models/dto"
	"testProject/middleware"
)

// CreateSongController is a controller for creating songs
type CreateSongController struct {
	DB *gorm.DB
}

// CreateSong godoc
// @Summary Create a new song
// @Description Create a new song with the provided details
// @ID create-song
// @Tags Songs
// @Accept  json
// @Produce  json
// @Param song body dto.SongDTO true "Song property"
// @Success 200 {object} dto.ResponseOK
// @Failure 400 {object} dto.ErrorBadRequestResponse
// @Failure 500 {object} dto.ErrorInternalServerResponse
// @Router /api/v1/create [get]
func (controller CreateSongController) CreateSong(ctx *gin.Context) {
	var songDTO dto.SongDTO

	if err := ctx.ShouldBindJSON(&songDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dt := songDTO.ReleaseDate

	nullTime, err := middleware.StringToNullTime(dt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Преобразование DTO в модель песню
	song := &models.Song{
		Group:       songDTO.Group,
		Song:        songDTO.Song,
		ReleaseDate: nullTime,
		Link:        songDTO.Link,
		Text:        songDTO.Text,
	}

	// Сохранение песню в базе данных
	if err := controller.DB.Create(&song).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Отправка ответа клиенту
	ctx.JSON(http.StatusOK, gin.H{
		"success": "true",
	})
}
