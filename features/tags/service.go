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

func GetTag(id string) (Tag, error) {
	var found = Tag{}

	stmt := `SELECT * FROM tags WHERE id = $1`

	err := db.DB.Get(&found, stmt, id)
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

	stmt := `INSERT INTO tags (id, name, key) VALUES (:id, :name, :key) RETURNING *`
	rows, er := tx.NamedQuery(stmt, tag)
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

	stmt := "UPDATE tags SET name = $2, key = $3 WHERE id = $1 RETURNING *"
	err := db.DB.QueryRow(stmt, id, tag.Name, tag.Key).Scan(&updatedTag.ID, &updatedTag.Name, &updatedTag.Key)

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

	stmt := `DELETE FROM tags WHERE id = $1`
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
