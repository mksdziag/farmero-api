package categories

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mksdziag/farmer-api/db"
)

func GetCategories() ([]Category, error) {
	var categories = make([]Category, 0)

	err := db.DB.Select(&categories, "SELECT * FROM categories")
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func GetCategory(id string) (Category, error) {
	var found = Category{}

	stmt := `SELECT * FROM categories WHERE id = $1`

	err := db.DB.Get(&found, stmt, id)
	if err != nil {
		return Category{}, err
	}

	return found, nil
}

func CreateCategory(category Category) (Category, error) {
	category.ID = uuid.New()

	tx, err := db.DB.Beginx()
	if err != nil {
		return Category{}, err
	}

	defer tx.Rollback()

	stmt := `INSERT INTO categories (id, name, description, key) VALUES (:id, :name, :description, :key) RETURNING *`
	rows, er := tx.NamedQuery(stmt, category)
	if err != nil {
		return Category{}, er
	}

	if err != nil {
		return Category{}, err
	}

	var inserted Category
	if rows.Next() {
		err := rows.StructScan(&inserted)
		if err != nil {
			return Category{}, err
		}
	}

	if err := tx.Commit(); err != nil {
		return Category{}, err
	}

	return inserted, nil
}

func UpdateCategory(id string, category Category) (Category, error) {
	var updatedCategory Category

	stmt := "UPDATE categories SET name = $2, description = $3, key = $4 WHERE id = $1 RETURNING *"
	err := db.DB.QueryRow(stmt, id, category.Name, category.Description, category.Key).Scan(&updatedCategory.ID, &updatedCategory.Name, &updatedCategory.Description, &updatedCategory.Key)

	if err != nil {
		return Category{}, err
	}

	return updatedCategory, nil
}

func DeleteCategory(id string) error {
	tx, err := db.DB.Beginx()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	stmt := `DELETE FROM categories WHERE id = $1`
	res, err := tx.Exec(stmt, id)

	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}

	err = tx.Commit()

	return err
}
