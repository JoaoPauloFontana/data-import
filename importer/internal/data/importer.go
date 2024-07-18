package data

import (
	"log"

	"github.com/JoaoPauloFontana/data-import/importer/internal/db"
	"github.com/xuri/excelize/v2"
)

func ImportData(filePath string) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal("Failed to open the Excel file:", err)
	} else {
		log.Println("Excel file opened successfully.")
	}

	rows, err := f.GetRows("Planilha1")
	if err != nil {
		log.Fatal("Failed to get rows from the sheet:", err)
	} else {
		log.Println("Rows fetched successfully.")
	}

	for _, row := range rows[1:] {
		_, err := db.DB.Exec(`
	      INSERT INTO records (
	          partner_id, partner_name, customer_id, customer_name, customer_domain_name, customer_country,
	          mpn_id, tier2_mpn_id, invoice_number, product_id, sku_id, availability_id, sku_name,
	          product_name, publisher_name, publisher_id, subscription_description, subscription_id,
	          charge_start_date, charge_end_date, usage_date, meter_type, meter_category, meter_id,
	          meter_sub_category, meter_name, meter_region, unit, resource_location, consumed_service,
	          resource_group, resource_uri, charge_type, unit_price, quantity, unit_type,
	          billing_pre_tax_total, billing_currency, pricing_pre_tax_total, pricing_currency,
	          service_info1, service_info2, tags, additional_info, effective_unit_price,
	          pct_to_bc_exchange_rate, pct_to_bc_exchange_rate_date, entitlement_id, entitlement_description,
	          partner_earned_credit_percentage, credit_percentage, credit_type, benefit_order_id,
	          benefit_id, benefit_type
	      ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55)`,
			row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8], row[9],
			row[10], row[11], row[12], row[13], row[14], row[15], row[16], row[17], row[18], row[19],
			row[20], row[21], row[22], row[23], row[24], row[25], row[26], row[27], row[28], row[29],
			row[30], row[31], row[32], row[33], row[34], row[35], row[36], row[37], row[38], row[39],
			row[40], row[41], row[42], row[43], row[44], row[45], row[46], row[47], row[48], row[49],
			row[50], row[51], " ", " ", " ",
		)

		if err != nil {
			log.Fatal("Failed to insert data:", err)
		} else {
			log.Println("Data inserted successfully.")
		}
	}
}
