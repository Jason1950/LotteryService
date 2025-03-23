package service

import (
	"LotteryService/internal/model"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/redis/go-redis/v9"
)

type AuthService struct {
	db        *sql.DB
	redis     *redis.Client
	jwtSecret []byte
}

func NewAuthService(db *sql.DB, redisClient *redis.Client) *AuthService {
	return &AuthService{
		db:        db,
		redis:     redisClient,
		jwtSecret: []byte("your-secret-key"), // 建議使用環境變量存儲
	}
}

type Claims struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func (s *AuthService) generateToken(user *model.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserId:   user.UserId,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "lottery-service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecret)
}

func (s *AuthService) storeUserToken(ctx context.Context, userId int64, token string) error {
	// 存儲用戶的 token
	key := fmt.Sprintf("user:%d:token", userId)
	// 設置 24 小時過期
	return s.redis.Set(ctx, key, token, 24*time.Hour).Err()
}

func (s *AuthService) Login(req *model.LoginRequest) (*model.LoginResponse, error) {
	ctx := context.Background()
	var user model.User

	// 查詢用戶
	err := s.db.QueryRow("SELECT username, password, balance, userid FROM custdb WHERE username = ?",
		req.Username).Scan(&user.Username, &user.Password, &user.Balance, &user.UserId)

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

	// 生成 JWT Token
	token, err := s.generateToken(&user)
	if err != nil {
		return nil, errors.New("生成token失敗")
	}

	// 將 token 存入 Redis
	err = s.storeUserToken(ctx, user.UserId, token)
	if err != nil {
		// 這裡我們只記錄錯誤，不影響登錄流程
		fmt.Printf("存儲token到Redis失敗: %v\n", err)
	}

	return &model.LoginResponse{
		Balance:  user.Balance,
		UserId:   user.UserId,
		Username: user.Username,
		Token:    token,
	}, nil
}

// 檢查 token 是否在黑名單中
func (s *AuthService) IsTokenBlacklisted(ctx context.Context, token string) bool {
	key := fmt.Sprintf("blacklist:%s", token)
	exists, _ := s.redis.Exists(ctx, key).Result()
	return exists > 0
}

// 將 token 加入黑名單
func (s *AuthService) BlacklistToken(ctx context.Context, token string, expiration time.Duration) error {
	key := fmt.Sprintf("blacklist:%s", token)
	return s.redis.Set(ctx, key, "1", expiration).Err()
}
