package dto

import "testProject/app/models"

type SongDTO struct {
	Group       string `json:"group" binding:"required" example:"Young Rus"`
	Song        string `json:"song" binding:"required" example:"NeonBones"`
	Text        string `json:"text" binding:"required" example:"I'm a sad guy, a humid day\n the humidity of my line"`
	Link        string `json:"link" binding:"required" example:"example.com"`
	ReleaseDate string `json:"releaseDate" binding:"required" example:"16.07.2006"`
}

type UpdateSongDTO struct {
	UpdateIdSong      string `json:"id" binding:"required" example:"1"`
	UpdateGroup       string `json:"group,omitempty" example:"Young Rus"`
	UpdateSong        string `json:"song,omitempty" example:"Fresh Soft"`
	UpdateReleaseDate string `json:"releaseDate,omitempty" example:"16.07.2026"`
	UpdateText        string `json:"text,omitempty" example:"I'm a sad guy, a humid day\n the humidity of my line\n the humidity of my line"`
	UpdateLink        string `json:"link,omitempty" example:""`
}

type DeleteSongDTO struct {
	DeleteIdSong int `json:"id" binding:"required" example:"1"`
}

type QueryTextSongDTO struct {
	IdSong string `json:"id" binding:"required" example:"1"`
}

type QueryInfoSongDTO struct {
	QueryGroup string `form:"group" binding:"required" example:"young Rus"`
	QuerySong  string `form:"song" binding:"required" example:"Neno Bones"`
}

type QueryPlayListDTO struct {
	QueryGroup       string `form:"group"  example:"young Rus"`
	QuerySong        string `form:"song"  example:"Neno Bones"`
	QueryReleaseDate string `form:"releaseDate"  example:"16.07.2006"`
	QueryText        string `json:"text" example:"I'm a sad guy, a humid day\n the humidity of my line"`
	QueryLink        string `json:"link" example:"example.com"`
	QueryPage        string `json:"page" example:"2" default:"1"`
	QueryLimit       string `json:"limit" example:"10" default:"10"`
}

type ResponseInfoSongDTO struct {
	Text        string `json:"text" example:"I'm a sad guy, a humid day\n the humidity of my line"`
	Link        string `json:"link" example:"example.com"`
	ReleaseDate string `json:"releaseDate" example:"16.07.2006"`
}

type ResponsePlayListSongDTO struct {
	Total int64         `json:"total" example:"1"`
	Page  int           `json:"page" example:"3"`
	Limit int           `json:"limit" example:"15"`
	Songs []models.Song `json:"data"`
}

type ResponseTextSongDTO struct {
	Text      string `json:"text" example:"I'm a sad guy, a humid day"`
	Part      int    `json:"part" example:"1"`
	TotalPart int    `json:"totalPart" example:"12"`
}
