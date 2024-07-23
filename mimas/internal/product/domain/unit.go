package domain

import "errors"

var (
	ErrUnitNameEmpty = errors.New("unit name cannot be empty")
)

type Unit struct {
	ID   int
	Name string
}

func NewUnit(name string) (*Unit, error) {
	unit := &Unit{
		Name: name,
	}

	if err := unit.Validate(); err != nil {
		return nil, err
	}

	return unit, nil
}

func (u *Unit) Validate() error {
	if u.Name == "" {
		return ErrUnitNameEmpty
	}

	if len(u.Name) < 2 || len(u.Name) > 2 {
		return ErrUnitNameEmpty
	}

	return nil
}
