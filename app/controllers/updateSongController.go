package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"testProject/app/models"
	"testProject/app/models/dto"
	"testProject/middleware"
)

type UpdateSongController struct {
	DB *gorm.DB
}

// UpdateSong godoc
// @Summary Update a song
// @Description Update the details of a song based on the provided information
// @Tags Songs
// @Accept json
// @Produce json
// @Param song body dto.UpdateSongDTO true "Song property"
// @Success 200 {object} dto.ResponseOK
// @Failure 400 {object} dto.ErrorBadRequestResponse
// @Failure 404 {object} dto.ErrorNotFoundResponse
// @Failure 500 {object} dto.ErrorInternalServerResponse
// @Router /songs/update [post]
func (controller UpdateSongController) UpdateSong(ctx *gin.Context) {
	var updatedSong dto.UpdateSongDTO
	if err := ctx.ShouldBindJSON(&updatedSong); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var song models.Song
	if err := controller.DB.First(&song, updatedSong.UpdateIdSong).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
		return
	}

	if updatedSong.UpdateGroup != "" {
		song.Group = updatedSong.UpdateGroup
	}
	if updatedSong.UpdateSong != "" {
		song.Song = updatedSong.UpdateSong
	}
	if updatedSong.UpdateReleaseDate != "" {
		nullTime, err := middleware.StringToNullTime(updatedSong.UpdateReleaseDate)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		song.ReleaseDate = nullTime
	}
	if updatedSong.UpdateText != "" {
		song.Text = updatedSong.UpdateText
	}
	if updatedSong.UpdateLink != "" {
		song.Link = updatedSong.UpdateLink
	}

	if err := controller.DB.Save(&song).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": "true",
	})
}
