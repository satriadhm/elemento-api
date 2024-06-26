package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Modul struct {
	gorm.Model
	ModulID     uuid.UUID `gorm:"column:id;type:char(36);primary_key;"`
	Title       string    `gorm:"column:nama_modul;type:varchar(255);"`
	Subtitle    string    `gorm:"column:subnama_modul;type:varchar(255);"`
	Image       string    `gorm:"column:image;type:varchar(255);"`
	ImageUrl    string    `gorm:"column:image_url;type:varchar(255);"`
	Progress    bool      `gorm:"column:progress;type:boolean;"`
	YoutubeLink string    `gorm:"column:youtube_link;type:varchar(255);"`
	Babs        []Bab     `gorm:"foreignkey:ModulID"`
}

type Bab struct {
	gorm.Model
	ModulID     uuid.UUID `gorm:"column:bab_modul;type:char(36);"`
	Title       string    `gorm:"column:nama_bab;type:varchar(255);"`
	Description string    `gorm:"column:deskripsi;type:text"`
	Task        string    `gorm:"column:task;type:text"`
}

func (u *Modul) TableName() string {
	return "moduls"
}

func (u *Bab) TableName() string {
	return "babs"
}

// Path: app/models/modul.go
