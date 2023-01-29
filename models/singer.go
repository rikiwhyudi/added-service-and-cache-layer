package models

type Singer struct {
	ID        int      `json:"id" gorm:"primary_key:auto_increment"`
	Name      string   `json:"name" gorm:"type: varchar(255)"`
	Thumbnail string   `json:"thumbnail" gorm:"type: varchar(255)"`
	Old       int      `json:"old" gorm:"type: varchar(255)"`
	Category  string   `json:"category" gorm:"type: varchar(255)"`
	Musics    []*Music `json:"musics" gorm:"many2many:music_singers"` //many to many

}

type SingerResponse struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Thumbnail string   `json:"thumbnail"`
	Old       string   `json:"old"`
	Category  string   `json:"category"`
	Musics    []*Music `json:"musics"`
}

func (SingerResponse) TableName() string {
	return "singers"
}
