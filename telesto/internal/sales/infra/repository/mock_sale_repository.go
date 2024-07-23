package repository

import (
	"github.com/jonattasmoraes/telesto/internal/sales/domain"
	"github.com/stretchr/testify/mock"
)

type MockSaleRepository struct {
	mock.Mock
}

func (s *MockSaleRepository) Save(sale *domain.Sale) error {
	args := s.Called(sale)
	return args.Error(0)
}