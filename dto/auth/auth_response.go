package authdto

type RegisterResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

type LoginResponse struct {
	Id    int    `json:"-"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
	Role  string `json:"role"`
}

type CheckAutResponse struct {
	Id    int    `json:"-"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}
