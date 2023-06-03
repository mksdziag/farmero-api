package articles

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mksdziag/farmer-api/db"
	"github.com/mksdziag/farmer-api/features/categories"
	"github.com/mksdziag/farmer-api/features/tags"
)

func GetArticlesByCategory(category string) ([]Article, error) {
	var articles = make([]Article, 0)

	db.DB.Select(&articles, "SELECT * FROM articles WHERE category = $1", category)

	return articles, nil
}

func GetArticles() ([]Article, error) {
	var articles = make([]Article, 0)

	err := db.DB.Select(&articles, "SELECT * FROM articles")
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func GetArticle(id string) (Article, error) {
	var found = Article{}

	query := `SELECT * FROM articles WHERE id = $1`

	err := db.DB.Get(&found, query, id)
	if err != nil {
		return Article{}, err
	}

	found.Categories, err = categories.GetCategoriesByArticle(id)
	if err != nil {
		return Article{}, err
	}

	found.Tags, err = tags.GetTagsByArticle(id)
	if err != nil {
		return Article{}, err
	}

	return found, nil
}

func CreateArticle(article Article) (Article, error) {
	article.ID = uuid.New()

	tx, err := db.DB.Beginx()
	if err != nil {
		return Article{}, err
	}

	defer tx.Rollback()

	query := `INSERT INTO articles (id, title, description, content, cover) VALUES (:id, :title, :description, :content, :cover) RETURNING *`
	rows, er := tx.NamedQuery(query, article)
	if err != nil {
		return Article{}, er
	}

	if err != nil {
		return Article{}, err
	}

	var inserted Article
	if rows.Next() {
		err = rows.StructScan(&inserted)
		if err != nil {
			return Article{}, err
		}
	}

	if err := tx.Commit(); err != nil {
		return Article{}, err
	}

	return inserted, nil
}

func UpdateArticle(id string, article Article) (Article, error) {
	var updatedArticle Article

	query := "UPDATE articles SET title = $2, description = $3, content = $4, cover = $5 WHERE id = $1 RETURNING *"
	err := db.DB.QueryRow(query, id, article.Title, article.Description, article.Content, article.Cover).Scan(&updatedArticle.ID, &updatedArticle.Title, &updatedArticle.Description, &updatedArticle.Content, &updatedArticle.Cover)

	if err != nil {
		return Article{}, err
	}

	return updatedArticle, nil
}

func DeleteArticle(id string) error {
	sqlStatement := `DELETE FROM articles WHERE id = $1`
	res, err := db.DB.Exec(sqlStatement, id)

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

	return nil
}
