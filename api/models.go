package api

import "time"

// backend models

type User struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type Property struct {
	ID          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Price       float64   `json:"price" db:"price"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UserID      int       `json:"user_id" db:"user_id"`
}

type JWTToken struct {
	Token string `json:"token"`
}
