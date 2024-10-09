package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"testProject/app/models"
	"testProject/app/models/dto"
	"testProject/middleware"
)

type GetDetailInfoSongController struct {
	DB *gorm.DB
}

// GetDetailInfoSong godoc
// @Summary Get detailed information about a song
// @Description Retrieve detailed information about a song based on the provided query parameters
// @Tags Songs
// @Accept json
// @Produce json
// @Param group query string true "Group of the song"
// @Param song query string true "Name of the song"
// @Success 200 {object} dto.ResponseInfoSongDTO
// @Failure 400 {object} dto.ErrorBadRequestResponse
// @Failure 500 {object} dto.ErrorInternalServerResponse
// @Router /api/v1/info [get]
func (controller GetDetailInfoSongController) GetDetailInfoSong(ctx *gin.Context) {
	var Query dto.QueryInfoSongDTO
	if err := ctx.ShouldBindQuery(&Query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var songs models.Song
	if err := controller.DB.Where(models.Song{Group: Query.QueryGroup, Song: Query.QuerySong}).First(&songs).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	songInfo := dto.ResponseInfoSongDTO{
		Link:        songs.Link,
		Text:        songs.Text,
		ReleaseDate: middleware.NullTimeToString(songs.ReleaseDate, nil),
	}

	ctx.JSON(http.StatusOK, songInfo)
}
