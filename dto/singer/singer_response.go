package singerdto

import musicdto "backend-api/dto/music"

type SingerResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Thumbnail   string `json:"thumbnail"`
	Old         int    `json:"old"`
	Category    string `json:"category"`
	StartCareer string `json:"start_career"`
}

type AllSingerResponse struct {
	ID          int                         `json:"id"`
	Name        string                      `json:"name"`
	Thumbnail   string                      `json:"thumbnail"`
	Old         int                         `json:"old"`
	Category    string                      `json:"category"`
	StartCareer string                      `json:"start_career"`
	Musics      []musicdto.AllMusicResponse `json:"musics"`
}
