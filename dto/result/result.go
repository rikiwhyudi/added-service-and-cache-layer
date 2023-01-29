package dto

type SuccessResult struct {
	Status int         `json:"status"`
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

type ErrorResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
