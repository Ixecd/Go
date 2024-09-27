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