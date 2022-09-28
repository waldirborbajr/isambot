package repository

import (
	"database/sql"
	"errors"
	"localhost/isam/pkg/entity"

	"github.com/mattn/go-sqlite3"
)

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("row not exists")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

// TODO
type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

func (r *SQLiteRepository) Migrate() error {
	query := `
    CREATE TABLE IF NOT EXISTS shoplist(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        description TEXT NOT NULL UNIQUE,
        quantity INTEGER NOT NULL
    );
    `

	_, err := r.db.Exec(query)
	return err
}

func (r *SQLiteRepository) Create(shoplist entity.ShopList) (*entity.ShopList, error) {
	res, err := r.db.Exec("INSERT INTO shoplist(description, quantity) values(?,?)", shoplist.Description, shoplist.Quantity)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
				return nil, ErrDuplicate
			}
		}
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	shoplist.ID = id

	return &shoplist, nil
}

func (r *SQLiteRepository) All() ([]entity.ShopList, error) {
	rows, err := r.db.Query("SELECT * FROM shoplist")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []entity.ShopList
	for rows.Next() {
		var shoplist entity.ShopList
		if err := rows.Scan(&shoplist.ID, &shoplist.Description, &shoplist.Quantity); err != nil {
			return nil, err
		}
		all = append(all, shoplist)
	}
	return all, nil
}

func (r *SQLiteRepository) GetByName(description string) (*entity.ShopList, error) {
	row := r.db.QueryRow("SELECT * FROM shoplist WHERE description = ?", description)

	var shoplist entity.ShopList
	if err := row.Scan(&shoplist.ID, &shoplist.Description, &shoplist.Quantity); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &shoplist, nil
}

func (r *SQLiteRepository) Update(id int64, updated entity.ShopList) (*entity.ShopList, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	res, err := r.db.Exec("UPDATE shoplist SET description = ?, quantity = ? WHERE id = ?", updated.Description, updated.Quantity, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, ErrUpdateFailed
	}

	return &updated, nil
}

func (r *SQLiteRepository) Delete(id int64) error {
	res, err := r.db.Exec("DELETE FROM shoplist WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrDeleteFailed
	}

	return err
}
