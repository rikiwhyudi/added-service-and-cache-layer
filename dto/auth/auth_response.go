package authdto

type RegisterResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Status   string `json:"status"`
}

type LoginResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Token  string `json:"token"`
	Status string `json:"status"`
}

type CheckAutResponse struct {
	Id     int    `json:"-"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
