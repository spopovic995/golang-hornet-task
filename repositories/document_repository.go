package repositories

import (
	"database/sql"
	model "golang-hornet-task/models"
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
