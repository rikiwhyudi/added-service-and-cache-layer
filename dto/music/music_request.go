package musicdto

import "time"

type MusicRequest struct {
	Title string `json:"title" form:"title" validate:"required"`
	// Thumbnail string    `json:"thumbnail" form:"thumbnail" validate:"required"`
	Year     time.Time `json:"year" form:"year" validate:"required"`
	SingerID int       `json:"singer_id" form:"singer_id" validate:"required"`
	// MusicFile string    `json:"music_file" form:"music_file" validate:"required"`
}

type UpdatedMusicRequest struct {
	Title string `json:"title" form:"title"`
	// Thumbnail string    `json:"thumbnail" form:"thumbnail"`
	Year time.Time `json:"year" form:"year"`
	// SingerID int       `json:"singer_id" form:"singer_id"`
	// MusicFile string    `json:"music_file" form:"music_file"`
}
