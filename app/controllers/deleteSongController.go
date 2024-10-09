package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"testProject/app/models"
	"testProject/app/models/dto"
)

type DeleteSongController struct {
	DB *gorm.DB
}

// DeleteSong @Summary Delete song
// @Description Delete song id
// @ID delete-song
// @Tags Songs
// @Accept  json
// @Produce  json
// @Param song body dto.DeleteSongDTO true "id delete song"
// @Success 200 {object} dto.ResponseOK
// @Failure 400 {object} dto.ErrorBadRequestResponse
// @Failure 404 {object} dto.ErrorNotFoundResponse
// @Router /api/v1/delete [get]
func (controller DeleteSongController) DeleteSong(ctx *gin.Context) {

	var deleteSong dto.DeleteSongDTO
	if err := ctx.ShouldBindJSON(&deleteSong); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var song models.Song
	if err := controller.DB.First(&song, deleteSong.DeleteIdSong).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
		return
	}
	controller.DB.Delete(&song)

	ctx.JSON(http.StatusOK, gin.H{
		"success": "true",
	})
}
