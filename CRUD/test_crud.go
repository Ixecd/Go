package main

import (
	"context"
	"database/sql"
	"log"
	"time"
)

var (
	ctx context.Context
	db *sql.DB
)

// CRUD
// 1. DATABASE/SQL
func straightforward() {
	id := 123
	var username string
	var created time.Time
	// 直接使用database/sql,速度非常快,也很直观
	// 但是非常死板,很无聊,容易出错,错误将仅在运行时显示
	err := db.QueryRowContext(ctx, "select username, created_at from users where id = ?", id).Scan(&username, &created)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("no user with id %d\n", id)
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:
		log.Printf("username is %q, account created on %s\n", username, created)
	}
}

// 2. GORM
// Golang的高级对象关系映射库
// 所有的CRUD操作都已经实现好了,生产代码很少
