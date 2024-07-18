package main

import (
	"log"

	"github.com/JoaoPauloFontana/data-import/api/internal/db"
	"github.com/JoaoPauloFontana/data-import/api/internal/routes"
)

func main() {
	db.Connect()

	e := routes.Init()

	log.Println("Starting server on :8080")
	e.Logger.Fatal(e.Start(":8080"))
}
