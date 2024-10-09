package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
	"testProject/app/models"
	"testProject/app/models/dto"
)

type GetTrackTextController struct {
	DB *gorm.DB
}

// GetTrackText godoc
// @Summary Get a specific part of the track text
// @Description Retrieve a specific part of the track text based on the provided page number and song ID
// @Tags Songs
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param body body dto.QueryTextSongDTO true "Song ID"
// @Success 200 {object} dto.ResponseTextSongDTO
// @Failure 400 {object} dto.ErrorBadRequestResponse
// @Failure 404 {object} dto.ErrorNotFoundResponse
// @Router /api/v1/getTrackText [get]
func (controller GetTrackTextController) GetTrackText(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))

	var deleteSong dto.QueryTextSongDTO
	if err := ctx.ShouldBindJSON(&deleteSong); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var song models.Song
	if err := controller.DB.First(&song, deleteSong.IdSong).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
		return
	}

	lines := strings.Split(song.Text, "\n")
	if len(lines) < page {
		page = len(lines)
		page--
	}

	ctx.JSON(http.StatusOK, dto.ResponseTextSongDTO{
		Text:      lines[page],
		Part:      page,
		TotalPart: len(lines) - 1,
	})
}
