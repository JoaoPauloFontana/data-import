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

	log.Println("Connecting to database...")

	var err error
	DB, err = sqlx.Connect("postgres", "user="+dbUser+" dbname="+dbName+" sslmode=disable password="+dbPassword+" host="+dbHost)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Connected to database")
	}
}

func Init() {
	schema := `
    CREATE TABLE IF NOT EXISTS records (
        id SERIAL PRIMARY KEY,
        partner_id TEXT,
        partner_name TEXT,
        customer_id TEXT,
        customer_name TEXT,
        customer_domain_name TEXT,
        customer_country TEXT,
        mpn_id TEXT,
        tier2_mpn_id TEXT,
        invoice_number TEXT,
        product_id TEXT,
        sku_id TEXT,
        availability_id TEXT,
        sku_name TEXT,
        product_name TEXT,
        publisher_name TEXT,
        publisher_id TEXT,
        subscription_description TEXT,
        subscription_id TEXT,
        charge_start_date DATE,
        charge_end_date DATE,
        usage_date DATE,
        meter_type TEXT,
        meter_category TEXT,
        meter_id TEXT,
        meter_sub_category TEXT,
        meter_name TEXT,
        meter_region TEXT,
        unit TEXT,
        resource_location TEXT,
        consumed_service TEXT,
        resource_group TEXT,
        resource_uri TEXT,
        charge_type TEXT,
        unit_price NUMERIC,
        quantity NUMERIC,
        unit_type TEXT,
        billing_pre_tax_total NUMERIC,
        billing_currency TEXT,
        pricing_pre_tax_total NUMERIC,
        pricing_currency TEXT,
        service_info1 TEXT,
        service_info2 TEXT,
        tags TEXT,
        additional_info TEXT,
        effective_unit_price NUMERIC,
        pct_to_bc_exchange_rate NUMERIC,
        pct_to_bc_exchange_rate_date DATE,
        entitlement_id TEXT,
        entitlement_description TEXT,
        partner_earned_credit_percentage NUMERIC,
        credit_percentage NUMERIC,
        credit_type TEXT,
        benefit_order_id TEXT,
        benefit_id TEXT,
        benefit_type TEXT
    );`
	DB.MustExec(schema)
}
