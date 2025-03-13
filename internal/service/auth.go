package service

import (
	"LotteryService/internal/model"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Login(req *model.LoginRequest) (*model.LoginResponse, error) {
	// TODO: 實現實際的登入邏輯
	// 這裡只是示例返回
	return &model.LoginResponse{
		Token: "dummy-token",
	}, nil
}
