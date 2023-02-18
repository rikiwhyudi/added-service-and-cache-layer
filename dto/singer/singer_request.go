package singerdto

import "time"

type SingerRequest struct {
	Name        string    `json:"name" form:"name" validate:"required"`
	Thumbnail   string    `json:"thumbnail" form:"thumbnail" validate:"required"`
	Old         int       `json:"old" form:"old" validate:"required"`
	Category    string    `json:"catergory" form:"category" validate:"required"`
	StartCareer time.Time `json:"start_career" form:"start_career" validate:"required"`
}

type UpdateSingerRequest struct {
	Name        string    `json:"name" form:"name"`
	Thumbnail   string    `json:"thumbnail" form:"thumbnail"`
	Old         int       `json:"old" form:"old"`
	Category    string    `json:"catergory" form:"category"`
	StartCareer time.Time `json:"start_career" form:"start_career"`
}
