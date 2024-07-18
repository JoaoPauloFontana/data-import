package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Connect() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	var err error
	DB, err = sqlx.Connect("postgres", "user="+dbUser+" dbname="+dbName+" sslmode=disable password="+dbPassword+" host="+dbHost)
	if err != nil {
		log.Fatalln(err)
	}
}
