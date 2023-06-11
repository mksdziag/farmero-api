package articles

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mksdziag/farmer-api/db"
	"github.com/mksdziag/farmer-api/features/categories"
	"github.com/mksdziag/farmer-api/features/tags"
)

func GetArticlesByCategoryId(id string) ([]Article, error) {
	var articles = make([]Article, 0)

	err := db.DB.Select(&articles, "SELECT * FROM articles WHERE id IN (SELECT article_id FROM articles_categories WHERE category_id = $1)", id)
	if err != nil {
		return nil, err
	}

	for idx := range articles {
		err = attachCategoriesToArticle(&articles[idx])
		if err != nil {
			return nil, err
		}

		err = attachTagsToArticle(&articles[idx])
		if err != nil {
			return nil, err
		}
	}

	return articles, nil
}

func GetArticlesByCategoryKey(key string) ([]Article, error) {
	category, err := categories.GetCategoryByKey(key)
	if err != nil {
		return nil, err
	}

	articles, err := GetArticlesByCategoryId(category.ID.String())
	if err != nil {
		return nil, err
	}

	for idx := range articles {
		err = attachCategoriesToArticle(&articles[idx])
		if err != nil {
			return nil, err
		}

		err = attachTagsToArticle(&articles[idx])
		if err != nil {
			return nil, err
		}
	}

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

	return found, nil
}

func attachCategoriesToArticle(article *Article) error {
	categories, err := categories.GetCategoriesByArticle(article.ID.String())
	article.Categories = categories
	fmt.Println(categories)
	fmt.Println(article.Categories)
	return err

}
func attachTagsToArticle(article *Article) error {
	tags, err := tags.GetTagsByArticle(article.ID.String())
	article.Tags = tags

	return err

}

func CreateArticle(article Article) (Article, error) {
	article.ID = uuid.New()

	tx, err := db.DB.Beginx()
	if err != nil {
		return Article{}, err
	}

	defer tx.Rollback()

	query := `INSERT INTO articles (id, title, description, content, cover) VALUES (:id, :title, :description, :content, :cover) RETURNING *`
	rows, err := tx.NamedQuery(query, article)
	if err != nil {
		return Article{}, err
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
	tx, err := db.DB.Beginx()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := `DELETE from articles_categories WHERE article_id = $1;`
	_, err = tx.Exec(query, id)
	if err != nil {
		return err
	}

	query = `DELETE from articles_tags WHERE article_id = $1;`
	_, err = tx.Exec(query, id)
	if err != nil {
		return err
	}

	query = `DELETE from articles WHERE id = $1`
	_, err = tx.Exec(query, id)
	if err != nil {
		return err
	}

	err = tx.Commit()

	return err
}
