package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
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
// @Router /songs/update [get]
func (controller UpdateSongController) UpdateSong(ctx *gin.Context) {
	log.Println("Starting UpdateSong function")

	var updatedSong dto.UpdateSongDTO
	if err := ctx.ShouldBindJSON(&updatedSong); err != nil {
		log.Printf("Error binding JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Received update request for song ID: %v", updatedSong.UpdateIdSong)

	var song models.Song
	if err := controller.DB.First(&song, updatedSong.UpdateIdSong).Error; err != nil {
		log.Printf("Song not found: %v", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
		return
	}

	if updatedSong.UpdateGroup != "" {
		log.Printf("Updating group to: %v", updatedSong.UpdateGroup)
		song.Group = updatedSong.UpdateGroup
	}
	if updatedSong.UpdateSong != "" {
		log.Printf("Updating song to: %v", updatedSong.UpdateSong)
		song.Song = updatedSong.UpdateSong
	}
	if updatedSong.UpdateReleaseDate != "" {
		nullTime, err := middleware.StringToNullTime(updatedSong.UpdateReleaseDate)
		if err != nil {
			log.Printf("Error converting release date: %v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Printf("Updating release date to: %v", nullTime)
		song.ReleaseDate = nullTime
	}
	if updatedSong.UpdateText != "" {
		log.Printf("Updating text to: %v", updatedSong.UpdateText)
		song.Text = updatedSong.UpdateText
	}
	if updatedSong.UpdateLink != "" {
		log.Printf("Updating link to: %v", updatedSong.UpdateLink)
		song.Link = updatedSong.UpdateLink
	}

	if err := controller.DB.Save(&song).Error; err != nil {
		log.Printf("Error saving song: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Song updated successfully: %v", song)
	ctx.JSON(http.StatusOK, gin.H{
		"success": "true",
	})
}
