package models

import "time"

type Singer struct {
	ID          int       `json:"id" gorm:"primary_key:auto_increment"`
	Name        string    `json:"name" gorm:"type: varchar(255)"`
	Thumbnail   string    `json:"thumbnail" gorm:"type: varchar(255)"`
	Old         int       `json:"old" gorm:"type: varchar(255)"`
	Category    string    `json:"category" gorm:"type: varchar(255)"`
	StartCareer time.Time `json:"start_career" gorm:"type: varchar(255)"`
	Musics      []*Music  `json:"musics" gorm:"many2many:music_singers"` //many to many

}

// type SingerResponse struct {
// 	ID          int       `json:"id"`
// 	Name        string    `json:"name"`
// 	Thumbnail   string    `json:"thumbnail"`
// 	Old         string    `json:"old"`
// 	Category    string    `json:"category"`
// 	StartCareer time.Time `json:"start_career"`
// 	Musics      []*Music  `json:"musics"`
// }

func (Singer) TableName() string {
	return "singers"
}
