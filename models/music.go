package models

import "time"

type Music struct {
	ID    int    `json:"id" gorm:"primary_key:auto_increment"`
	Title string `json:"title" gorm:"type: varchar(255)"`
	// Thumbnail string    `json:"thumbnail" gorm:"type: varchar(255)"`
	Year     time.Time `json:"year"`
	SingerID int       `json:"-" gorm:"foreignkey:SingerID"` // has many fields from singer
	Singer   Singer    `json:"singer"`
	// MusicFile string `json:"music_file" gorm:"type: varchar(255)"`
}

// type MusicResponse struct {
// 	ID         int    `json:"id"`
// 	Title      string `json:"title"`
// 	Year       int    `json:"year"`
// 	SingerName string `json:"singer_name"`
// 	MusicFile  string `json:"music_file"`
// }

func (Music) TableName() string {
	return "musics"
}
