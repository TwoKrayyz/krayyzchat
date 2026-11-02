package database

import (
	"os"
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Pool() error {
	if err := godotenv.Load("db.env"); err != nil {
		return err
	}

	driver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	data := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + name + " sslmode=" + sslmode

	db, err := sql.Open(driver, data)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	DB = db

	return nil
}