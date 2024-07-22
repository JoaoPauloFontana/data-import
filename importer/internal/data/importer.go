package data

import (
	"log"

	"github.com/JoaoPauloFontana/data-import/importer/internal/batch"
	"github.com/JoaoPauloFontana/data-import/importer/internal/db"
	"github.com/xuri/excelize/v2"
)

const batchSize = 1000

func ImportData(filePath string) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal("Failed to open the Excel file:", err)
	}
	log.Println("Excel file opened successfully.")

	sheetList := f.GetSheetList()
	if len(sheetList) == 0 {
		log.Fatal("No sheets found in the Excel file.")
	}
	sheetName := sheetList[0]
	log.Println("Sheet name fetched successfully:", sheetName)

	rows, err := f.Rows(sheetName)
	if err != nil {
		log.Fatal("Failed to get rows from the sheet:", err)
	}
	log.Println("Rows fetched successfully.")

	var records [][]interface{}
	isFirstRow := true

	for rows.Next() {
		row, err := rows.Columns()
		if err != nil {
			log.Fatal("Failed to get columns from row:", err)
		}
		if isFirstRow {
			isFirstRow = false
			continue
		}

		record := make([]interface{}, 55)
		for i, cell := range row {
			if i < len(record) {
				record[i] = cell
			} else {
				record[i] = ""
			}
		}
		records = append(records, record)

		if len(records) >= batchSize {
			if err := batch.InsertRecords(db.DB, records); err != nil {
				log.Fatal("Failed to insert data:", err)
			}
			log.Println("Batch inserted successfully.")
			records = nil
		}
	}

	if len(records) > 0 {
		if err := batch.InsertRecords(db.DB, records); err != nil {
			log.Fatal("Failed to insert data:", err)
		}
		log.Println("Final batch inserted successfully.")
	}

	log.Println("Data import completed successfully.")
}
