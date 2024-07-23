package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Helper function to create a sample product
func sampleProduct() SaleProduct {
	return SaleProduct{
		ProductID: "prod-123",
		Name:      "Sample Product",
		Unit:      "piece",
		Category:  "sample",
		Quantity:  2,
		Price:     100,
	}
}

func TestNewSale_Valid(t *testing.T) {
	// Create a new user with valid parameters
	_, err := NewSale("prod-123", []SaleProduct{sampleProduct()}, "credit")
	assert.NoError(t, err)
}


// Test NewSale function with valid data
func TestNewSale(t *testing.T) {
	userID := "user-123"
	products := []SaleProduct{sampleProduct()}
	paymentType := "cash"

	sale, err := NewSale(userID, products, paymentType)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if sale.UserID != userID {
		t.Errorf("Expected userID to be %v, got %v", userID, sale.UserID)
	}

	if len(sale.Products) != 1 {
		t.Errorf("Expected 1 product, got %d", len(sale.Products))
	}

	if sale.PaymentType != paymentType {
		t.Errorf("Expected paymentType to be %v, got %v", paymentType, sale.PaymentType)
	}

	expectedTotal := products[0].Price * products[0].Quantity
	if sale.Total != expectedTotal {
		t.Errorf("Expected total to be %d, got %d", expectedTotal, sale.Total)
	}
}

// Test NewSale with missing payment type
func TestNewSale_MissingPaymentType(t *testing.T) {
	_, err := NewSale("user-123", []SaleProduct{sampleProduct()}, "")
	if err != ErrRequiredPaymentType {
		t.Fatalf("Expected error %v, got %v", ErrRequiredPaymentType, err)
	}
}

// Test NewSale with invalid payment type
func TestNewSale_InvalidPaymentType(t *testing.T) {
	_, err := NewSale("user-123", []SaleProduct{sampleProduct()}, "invalid")
	if err != ErrInvalidPaymentType {
		t.Fatalf("Expected error %v, got %v", ErrInvalidPaymentType, err)
	}
}

// Test NewSale with no products
func TestNewSale_NoProducts(t *testing.T) {
	_, err := NewSale("user-123", []SaleProduct{}, "cash")
	if err != ErrProductsRequired {
		t.Fatalf("Expected error %v, got %v", ErrProductsRequired, err)
	}
}

// Test Validate function for missing payment type
func TestSale_Validate_MissingPaymentType(t *testing.T) {
	sale := &Sale{
		UserID: "user-123",
		Products: []SaleProduct{
			sampleProduct(),
		},
		PaymentType: "",
		Date:        time.Now(),
	}

	if err := sale.Validate(); err != ErrRequiredPaymentType {
		t.Fatalf("Expected error %v, got %v", ErrRequiredPaymentType, err)
	}
}

// Test Validate function for invalid payment type
func TestSale_Validate_InvalidPaymentType(t *testing.T) {
	sale := &Sale{
		UserID: "user-123",
		Products: []SaleProduct{
			sampleProduct(),
		},
		PaymentType: "invalid",
		Date:        time.Now(),
	}

	if err := sale.Validate(); err != ErrInvalidPaymentType {
		t.Fatalf("Expected error %v, got %v", ErrInvalidPaymentType, err)
	}
}

// Test Validate function for no products
func TestSale_Validate_NoProducts(t *testing.T) {
	sale := &Sale{
		UserID:      "user-123",
		Products:    []SaleProduct{},
		PaymentType: "cash",
		Date:        time.Now(),
	}

	if err := sale.Validate(); err != ErrProductsRequired {
		t.Fatalf("Expected error %v, got %v", ErrProductsRequired, err)
	}
}

