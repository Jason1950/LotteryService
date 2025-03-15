package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {
	// 替換以下配置為你的 MySQL 配置
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/custdb?parseTime=true")
	if err != nil {
		return nil, err
	}

	// 測試連接
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
