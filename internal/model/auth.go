package model

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Username string  `json:"username"`
	Balance  float64 `json:"balance"`
	UserId   float64 `json:"userid"`
}

type User struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
	UserId   float64 `json:"userid"`
}
