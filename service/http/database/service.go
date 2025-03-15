package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Client struct {
	DB *sql.DB
}

type Service interface {
	Client() (*Client, error)
}

type service struct {
	dbConn string
}

func NewService(connStr string) Service {
	return &service{dbConn: connStr}
}

func (s *service) Client() (*Client, error) {
	db, err := sql.Open("postgres", s.dbConn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	fmt.Println("Database connection established")
	return &Client{DB: db}, nil
}
