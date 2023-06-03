package tags

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mksdziag/farmer-api/db"
)

func GetTags() ([]Tag, error) {
	var tags = make([]Tag, 0)

	err := db.DB.Select(&tags, "SELECT * FROM tags")
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func GetTagsByArticle(id string) ([]Tag, error) {
	query := `
	SELECT t.id, t.name, t.key FROM articles_tags at
	JOIN tags t ON t.id = at.tag_id
	WHERE article_id = $1
	`
	rows, err := db.DB.Query(query, id)
	if err != nil {
		return nil, err
	}

	tagsList := make([]Tag, 0)
	if rows.Next() {
		var tag = Tag{}

		err := rows.Scan(&tag.ID, &tag.Name, &tag.Key)
		if err != nil {
			return nil, err
		}

		tagsList = append(tagsList, tag)
	}

	return tagsList, nil
}

func GetTag(id string) (Tag, error) {
	var found = Tag{}

	query := `SELECT * FROM tags WHERE id = $1`

	err := db.DB.Get(&found, query, id)
	if err != nil {
		return Tag{}, err
	}

	return found, nil
}

func CreateTag(tag Tag) (Tag, error) {
	tag.ID = uuid.New()

	tx, err := db.DB.Beginx()
	if err != nil {
		return Tag{}, err
	}

	defer tx.Rollback()

	query := `INSERT INTO tags (id, name, key) VALUES (:id, :name, :key) RETURNING *`
	rows, er := tx.NamedQuery(query, tag)
	if err != nil {
		return Tag{}, er
	}

	if err != nil {
		return Tag{}, err
	}

	var inserted Tag
	if rows.Next() {
		err := rows.StructScan(&inserted)
		if err != nil {
			return Tag{}, err
		}
	}

	if err := tx.Commit(); err != nil {
		return Tag{}, err
	}

	return inserted, nil
}

func UpdateTag(id string, tag Tag) (Tag, error) {
	var updatedTag Tag

	query := "UPDATE tags SET name = $2, key = $3 WHERE id = $1 RETURNING *"
	err := db.DB.QueryRow(query, id, tag.Name, tag.Key).Scan(&updatedTag.ID, &updatedTag.Name, &updatedTag.Key)

	if err != nil {
		return Tag{}, err
	}

	return updatedTag, nil
}

func DeleteTag(id string) error {
	tx, err := db.DB.Beginx()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := `DELETE FROM tags WHERE id = $1`
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
