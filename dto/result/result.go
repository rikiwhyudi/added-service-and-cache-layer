package dto

type SuccessResult struct {
	Status int         `json:"status"`
	Action string      `json:"action"`
	Data   interface{} `json:"user"`
}

type ArtistResult struct {
	Status int         `json:"status"`
	Action string      `json:"action"`
	Data   interface{} `json:"artist"`
}

type MusicResult struct {
	Status int         `json:"status"`
	Action string      `json:"action"`
	Data   interface{} `json:"music"`
}

type TransactionResult struct {
	Status int         `json:"status"`
	Action string      `json:"action"`
	Data   interface{} `json:"transaction"`
}

type ErrorResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
