package repositories

import (
	"database/sql"
	model "golang-hornet-task/models"
	"strings"
)

type DocumentRepository struct {
	DB *sql.DB
}

func NewDocumentRepository(db *sql.DB) *DocumentRepository {
	return &DocumentRepository{DB: db}
}

func (r *DocumentRepository) Save(doc model.Document) error {
	query := "INSERT INTO documents (name, description) VALUES (?, ?)"
	_, err := r.DB.Exec(query, doc.Name, doc.Description)
	return err
}

func (r *DocumentRepository) GetAll() ([]model.Document, error) {
	rows, err := r.DB.Query("SELECT id, name, description FROM documents")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var documents []model.Document
	for rows.Next() {
		var doc model.Document
		if err := rows.Scan(&doc.ID, &doc.Name, &doc.Description); err != nil {
			return nil, err
		}
		documents = append(documents, doc)
	}
	return documents, nil
}

func (r *DocumentRepository) GetByID(id int64) (model.Document, error) {
	var doc model.Document
	err := r.DB.QueryRow("SELECT id, name, description FROM documents WHERE id = ?", id).
		Scan(&doc.ID, &doc.Name, &doc.Description)
	return doc, err
}

func (r *DocumentRepository) DeleteById(id int64) bool {
	query := "DELETE FROM Documents WHERE id = ?"
	result, err := r.DB.Exec(query, id)
	if err != nil {
		return false
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false
	}

	return rowsAffected > 0

}

func (r *DocumentRepository) SearchByNameLogic(query string) ([]model.Document, error) {
	baseSQL := "SELECT id, name, description FROM documents WHERE"

	// this part here splits the query into an array of strings
	tokens := strings.Fields(query)

	var sqlParts []string
	var args []interface{}

	for i := 0; i < len(tokens); i++ {
		token := strings.ToUpper(tokens[i])
		switch token {
		case "AND", "OR":
			sqlParts = append(sqlParts, token)
		case "NOT":
			if i+1 < len(tokens) {
				sqlParts = append(sqlParts, "name NOT LIKE ?")
				args = append(args, "%"+tokens[i+1]+"%")
				i++
			}
		default:
			sqlParts = append(sqlParts, "name LIKE ?")
			args = append(args, "%"+tokens[i]+"%")
		}
	}

	sql := baseSQL + " " + strings.Join(sqlParts, " ")
	rows, err := r.DB.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []model.Document
	for rows.Next() {
		var doc model.Document
		if err := rows.Scan(&doc.ID, &doc.Name, &doc.Description); err != nil {
			return nil, err
		}
		results = append(results, doc)
	}
	return results, nil
}
