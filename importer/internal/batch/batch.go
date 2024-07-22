package batch

import "github.com/jmoiron/sqlx"

func InsertRecords(db *sqlx.DB, records [][]interface{}) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`
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
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, record := range records {
		_, err := stmt.Exec(record...)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
