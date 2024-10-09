package models

import (
	"database/sql"
	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	Group       string       `gorm:"size:100;not null;" json:"group"`
	Song        string       `gorm:"size:100;not null;" json:"song"`
	ReleaseDate sql.NullTime `gorm:"default:null" json:"releaseDate"`
	Text        string       `gorm:"not null;" json:"text"`
	Link        string       `gorm:"size:100;not null;" json:"link"`
}
