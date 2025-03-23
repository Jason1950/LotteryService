package model

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserId   int64   `json:"user_id"`
	Username string  `json:"username"`
	Balance  float64 `json:"balance"`
	Token    string  `json:"token"`
}

type User struct {
	UserId   int64   `json:"user_id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
}
