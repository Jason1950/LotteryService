package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {
	config := GetConfig()
	if config == nil {
		return nil, fmt.Errorf("配置未加載")
	}

	// 構建連接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DBName,
	)

	// 建立連接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("數據庫連接失敗: %v", err)
	}

	// 測試連接
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("數據庫連接測試失敗: %v", err)
	}

	// 設置連接池參數
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	return db, nil
}
