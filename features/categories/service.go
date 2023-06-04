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

	query := `SELECT * FROM categories WHERE id = $1`

	err := db.DB.Get(&found, query, id)
	if err != nil {
		return Category{}, err
	}

	return found, nil
}
func GetCategoryByKey(key string) (Category, error) {
	var found = Category{}

	query := `SELECT * FROM categories WHERE key = $1`

	err := db.DB.Get(&found, query, key)
	if err != nil {
		return Category{}, err
	}

	return found, nil
}

func GetCategoriesByArticle(articleID string) ([]Category, error) {
	query := `
			SELECT c.id, c.name, c.description, c.key
			FROM articles_categories ac
			JOIN categories c ON c.id = ac.category_id
			WHERE article_id = $1
	`
	rows, err := db.DB.Query(query, articleID)
	if err != nil {
		return nil, err
	}

	categories := make([]Category, 0)
	if rows.Next() {
		var category Category

		err := rows.Scan(&category.ID, &category.Name, &category.Description, &category.Key)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func CreateCategory(category Category) (Category, error) {
	category.ID = uuid.New()

	tx, err := db.DB.Beginx()
	if err != nil {
		return Category{}, err
	}

	defer tx.Rollback()

	query := `INSERT INTO categories (id, name, description, key) VALUES (:id, :name, :description, :key) RETURNING *`
	rows, er := tx.NamedQuery(query, category)
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

	query := "UPDATE categories SET name = $2, description = $3, key = $4 WHERE id = $1 RETURNING *"
	err := db.DB.QueryRow(query, id, category.Name, category.Description, category.Key).Scan(&updatedCategory.ID, &updatedCategory.Name, &updatedCategory.Description, &updatedCategory.Key)

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

	query := `DELETE FROM categories WHERE id = $1`
	res, err := tx.Exec(query, id)

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
