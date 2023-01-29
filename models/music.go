package models

type Music struct {
	ID        int    `json:"id" gorm:"primary_key:auto_increment"`
	Title     string `json:"title" gorm:"type: varchar(255)"`
	Year      int    `json:"year" gorm:"type: varchar(255)"`
	SingerID  int    `json:"singer_id" gorm:"foreignkey:SingerID"`  //many to many
	Singer    Singer `json:"singer" gorm:"many2many:music_singers"` //many to many relationships
	MusicFile string `json:"music_file" gorm:"type: varchar(255)"`
}

type MusicResponse struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Year       int    `json:"year"`
	SingerName string `json:"singer_name"`
	MusicFile  string `json:"music_file"`
}

func (MusicResponse) TableName() string {
	return "musics"
}
