package api

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Db connection for pg and db table creation func

var db *sqlx.DB

func InitDB() {
	var err error
	connStr := os.Getenv("DB")
	db, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}
	fmt.Println("Connected to database")
}

func RunMigrations() error {
	migrations := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS properties (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			description TEXT NOT NULL,
			price NUMERIC(10, 2) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			user_id INT NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		);`,
	}

	for _, migration := range migrations {
		_, err := db.Exec(migration)
		if err != nil {
			return fmt.Errorf("failed to execute migration: %v", err)
		}
	}

	return nil
}
