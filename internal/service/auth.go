package service

import (
	"LotteryService/internal/model"
	"database/sql"
	"errors"
)

type AuthService struct {
	db *sql.DB
}

func NewAuthService(db *sql.DB) *AuthService {
	return &AuthService{
		db: db,
	}
}

func (s *AuthService) Login(req *model.LoginRequest) (*model.LoginResponse, error) {
	var user model.User

	// 查詢用戶
	err := s.db.QueryRow("SELECT username, password, balance FROM custdb WHERE username = ?",
		req.Username).Scan(&user.Username, &user.Password, &user.Balance)

	if err == sql.ErrNoRows {
		return nil, errors.New("用戶不存在")
	}
	if err != nil {
		return nil, err
	}

	// 驗證密碼
	if user.Password != req.Password {
		return nil, errors.New("密碼錯誤")
	}

	return &model.LoginResponse{
		Balance: user.Balance,
	}, nil
}
