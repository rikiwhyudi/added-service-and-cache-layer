package models

import "time"

type Singer struct {
	ID          int       `json:"id" gorm:"primary_key:auto_increment"`
	Name        string    `json:"name" gorm:"type: varchar(255)"`
	Thumbnail   string    `json:"thumbnail" gorm:"type: varchar(255)"`
	Old         int       `json:"old"`
	Category    string    `json:"category" gorm:"type: varchar(255)"`
	StartCareer time.Time `json:"start_career"`
	Music       []Music   `json:"musics" gorm:"foreginKey:SingerID"` //has many

}

func (Singer) TableName() string {
	return "singers"
}
