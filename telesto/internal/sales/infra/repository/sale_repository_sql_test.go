package repository_test

import (
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/jonattasmoraes/telesto/internal/sales/domain"
	"github.com/jonattasmoraes/telesto/internal/sales/infra/repository"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestSaveSale(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := repository.NewSalesRepository(db, db)

	sale := &domain.Sale{
		SaleID:      "sale-123",
		UserID:      "user-123",
		Products:    []domain.SaleProduct{sampleProduct()},
		Total:       200,
		PaymentType: "cash",
		Date:        time.Now(),
	}

	err := repo.Save(sale)
	require.NoError(t, err)
}

func setupTestDB(t *testing.T) *sqlx.DB {
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to connect to SQLite in-memory database: %v", err)
	}

	// Create sales table
	_, err = db.Exec(`
	CREATE TABLE sales (
		id VARCHAR(255) PRIMARY KEY,
		user_id VARCHAR(255) NOT NULL,
		total INTEGER NOT NULL,
		payment_type VARCHAR(50) NOT NULL,
		date TIMESTAMP NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	)
	`)
	if err != nil {
		t.Fatalf("Failed to create sales table: %v", err)
	}

	// Create sale_products table
	_, err = db.Exec(`
	CREATE TABLE sale_products (
		id SERIAL PRIMARY KEY,
		sale_id VARCHAR(255) REFERENCES sales(id) ON DELETE CASCADE,
		product_id VARCHAR(255) NOT NULL,
		quantity INTEGER NOT NULL,
		price INTEGER NOT NULL
	)
	`)
	if err != nil {
		t.Fatalf("Failed to create sale_products table: %v", err)
	}

	return db
}

func sampleProduct() domain.SaleProduct {
	return domain.SaleProduct{
		ProductID: "prod-123",
		Name:      "Sample Product",
		Unit:      "piece",
		Category:  "sample",
		Quantity:  2,
		Price:     100,
	}
}