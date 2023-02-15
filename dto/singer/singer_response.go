package singerdto

import (
	"backend-api/models"
	"time"
)

type SingerResponse struct {
	ID          int             `json:"-"`
	Name        string          `json:"name"`
	Thumbnail   string          `json:"thumbnail"`
	Old         int             `json:"old"`
	Category    string          `json:"category"`
	StartCareer time.Time       `json:"start_career"`
	Musics      []*models.Music `json:"-"`
}
