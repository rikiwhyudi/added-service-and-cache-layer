package musicdto

type MusicResponse struct {
	ID         int         `json:"id"`
	Title      string      `json:"title"`
	Thumbnail  string      `json:"-"`
	Year       string      `json:"year"`
	SingerName interface{} `json:"singer"`
	MusicFile  string      `json:"-"`
}

type AllMusicResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Thumbnail string `json:"-"`
	Year      string `json:"year"`
	MusicFile string `json:"-"`
}
