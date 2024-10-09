package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"testProject/app/models"
	"testProject/app/models/dto"
	"testProject/middleware"
)

type GetPlayListController struct {
	DB *gorm.DB
}

// GetPlayList godoc
// @Summary Get a list of songs with pagination and filtering
// @Description Retrieve a list of songs with pagination and optional filtering by group, song, release date, text, and link
// @Tags Songs
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Param group query string false "Filter by group"
// @Param song query string false "Filter by song"
// @Param releaseDate query string false "Filter by release date"
// @Param text query string false "Filter by text"
// @Param link query string false "Filter by link"
// @Success 200 {object} dto.PlaylistResponseDto "List of songs"
// @Failure 500 {object} dto.ErrorInternalServerResponse
// @Router /api/v1/playlist [get]
func (controller GetPlayListController) GetPlayList(ctx *gin.Context) {
	var Query dto.QueryPlayListDTO
	if err := ctx.ShouldBindQuery(&Query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	query := controller.DB.Model(&models.Song{})

	if Query.QueryGroup != "" {
		query = query.Where(&models.Song{Group: Query.QueryGroup})
	}
	if Query.QuerySong != "" {
		query = query.Where(&models.Song{Song: Query.QuerySong})
	}
	if Query.QueryReleaseDate != "" {
		nullTime, err := middleware.StringToNullTime(Query.QueryReleaseDate)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		query = query.Where(&models.Song{ReleaseDate: nullTime})
	}
	if Query.QueryText != "" {
		query = query.Where(&models.Song{Text: Query.QueryText})
	}
	if Query.QueryLink != "" {
		query = query.Where(&models.Song{Link: Query.QueryLink})
	}

	var total int64
	query.Count(&total)

	var songs []models.Song
	if err := query.Offset(offset).Limit(limit).Find(&songs).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.ResponsePlayListSongDTO{
		Total: total,
		Page:  page,
		Limit: limit,
		Songs: songs,
	}

	ctx.JSON(http.StatusOK, response)
}
