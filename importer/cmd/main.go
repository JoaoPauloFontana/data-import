package main

import (
	"log"

	"github.com/JoaoPauloFontana/data-import/importer/internal/data"
	"github.com/JoaoPauloFontana/data-import/importer/internal/db"
)

func main() {
	db.Connect()
	db.Init()

	filePath := "./Reconfile_fornecedores.xlsx"
	data.ImportData(filePath)

	log.Println("Data import completed successfully.")
}
