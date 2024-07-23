package repository

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/jonattasmoraes/mimas/internal/product/domain"
)

var (
	ErrUnitNotFound = errors.New("unit not found, please enter a valid unit and try again")
	ErrAlreadyUsed  = errors.New("already in use, cannot be deleted")
)

type UnitSqlx struct {
	writer *sqlx.DB
	reader *sqlx.DB
}

func NewUnitSqlxRepository(writer, reader *sqlx.DB) domain.UnitRepository {
	return &UnitSqlx{writer: writer, reader: reader}
}

func (u *UnitSqlx) CreateUnit(unity *domain.Unit) error {
	query := `
	INSERT INTO unit (name)
	VALUES ($1)
	`

	_, err := u.writer.Exec(query, unity.Name)
	if err != nil {
		return err
	}

	return nil
}

func (u *UnitSqlx) ListUnits() ([]*domain.Unit, error) {
	query := `	
	SELECT id, name
	FROM unit
	`

	var units []*domain.Unit
	err := u.reader.Select(&units, query)
	if err != nil {
		return nil, err
	}

	return units, nil
}

func (u *UnitSqlx) FindUnitById(id string) (*domain.Unit, error) {
	query := `
	SELECT id, name
	FROM unit
	WHERE id = $1
	`

	var unit domain.Unit
	err := u.reader.QueryRow(query, id).Scan(&unit.ID, &unit.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUnitNotFound
		}
		return nil, err
	}

	return &unit, nil
}

func (r *UnitSqlx) DeleteUnit(id string) error {
	_, err := r.writer.Exec(`
	DELETE FROM unit
	WHERE id = $1`,
		id,
	)

	if err != nil {
		if strings.Contains(err.Error(), "violates foreign key constraint") {
			return ErrAlreadyUsed
		}
		return err
	}

	return err
}
