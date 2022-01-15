package mysql

import (
	"database/sql"
	"errors"

	"github.com/mapspiral/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (model *SnippetModel) Insert(title string, content string, expires string) (int, error) {
	statement := `INSERT INTO snippets (title, content, created, expires)
			VALUES (?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, errorInfo := model.DB.Exec(statement, title, content, expires)

	if errorInfo != nil {
		return 0, errorInfo
	}

	id, errorInfo := result.LastInsertId()

	if errorInfo != nil {
		return 0, errorInfo
	}

	return int(id), nil
}

func (model *SnippetModel) Get(id int) (*models.Snippet, error) {
	statement := `SELECT id, title, content, created, expires
				  FROM snippets 
				  WHERE expires > UTC_TIMESTAMP() AND id = ?`

	result := model.DB.QueryRow(statement, id)

	snippet := &models.Snippet{}

	errorInfo := result.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires)

	if errorInfo != nil {
		if errors.Is(errorInfo, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, errorInfo
		}
	}

	return snippet, nil
}

func (model *SnippetModel) Latest() ([]*models.Snippet, error) {
	statement := `SELECT id, title, content, created, expires
				  FROM snippets 
				  WHERE expires > UTC_TIMESTAMP()
				  ORDER BY CREATED
				  DESC LIMIT 10`

	rows, errorInfo := model.DB.Query(statement)

	if errorInfo != nil {
		return nil, errorInfo
	}

	defer rows.Close()

	snippets := []*models.Snippet{}

	for rows.Next() {
		snippet := &models.Snippet{}

		errorInfo := rows.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires)

		if errorInfo != nil {
			return nil, errorInfo
		}

		snippets = append(snippets, snippet)
	}

	if errorInfo = rows.Err(); errorInfo != nil {
		return nil, errorInfo
	}

	return snippets, nil
}
