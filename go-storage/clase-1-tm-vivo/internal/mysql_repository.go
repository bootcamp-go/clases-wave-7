package products

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

type MySQLRepository struct {
	Repository

	Database *sql.DB
}

func (repository *MySQLRepository) Get(id int) (product *Product, err error) {
	// Query product.
	query := `
		SELECT
			id, name, type, count, price 
		FROM products 
		WHERE 
			id = ?;
	`

	row := repository.Database.QueryRow(query, id)

	// Check if product exists.
	if row.Err() != nil {
		switch row.Err() {
		case sql.ErrNoRows:
			err = ErrNotFound
		default:
			err = ErrInternal
		}
		return
	}

	// Scan product.
	err = row.Scan(
		&product.ID,
		&product.Name,
		&product.Type,
		&product.Count,
		&product.Price,
	)
	if err != nil {
		err = ErrInternal
		return
	}

	// Everything is ok.
	return
}

func (repository *MySQLRepository) Store(product *Product) (err error) {
	// Insert product.
	statement, err := repository.Database.Prepare(`
		INSERT INTO products (
			name, type, count, price
		) VALUES (
			?, ?, ?, ?
		);
	`)
	if err != nil {
		err = ErrInternal
		return
	}
	defer statement.Close()

	result, err := statement.Exec(
		product.Name,
		product.Type,
		product.Count,
		product.Price,
	)
	if err != nil {
		driverErr, ok := err.(*mysql.MySQLError)
		if !ok {
			err = ErrInternal
			return
		}

		switch driverErr.Number {
		case 1062:
			err = ErrDuplicated
		default:
			err = ErrInternal
		}

		return
	}

	// Check if product was inserted.
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected != 1 {
		err = ErrInternal
		return
	}

	// Get product id.
	productID, err := result.LastInsertId()
	if err != nil {
		err = ErrInternal
		return
	}

	// Everything is ok.
	product.ID = int(productID)
	return
}

func (repository *MySQLRepository) Update(product *Product) (err error) {
	// Update product.
	statement, err := repository.Database.Prepare(`
		UPDATE products SET
			name = ?, type = ?, count = ?, price = ?
		WHERE
			id = ?;
	`)
	if err != nil {
		err = ErrInternal
		return
	}

	result, err := statement.Exec(
		product.Name,
		product.Type,
		product.Count,
		product.Price,
		product.ID,
	)
	if err != nil {
		// TODO: Validar según error code.
		err = ErrInternal
		return
	}

	// Check if product was updated.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		err = ErrInternal
		return
	}

	if rowsAffected != 1 {
		err = ErrNotFound
		return
	}

	// Everything is ok.
	return
}

func (repository *MySQLRepository) Delete(id int) (err error) {
	// Delete product.
	statement, err := repository.Database.Prepare(`
		DELETE FROM products
		WHERE
			id = ?;
	`)
	if err != nil {
		err = ErrInternal
		return
	}

	result, err := statement.Exec(id)
	if err != nil {
		// TODO: Validar según error code.
		err = ErrInternal
		return
	}

	// Check if product was deleted.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		err = ErrInternal
		return
	}

	if rowsAffected == 0 {
		err = ErrNotFound
		return
	}

	// Everything is ok.
	return
}
