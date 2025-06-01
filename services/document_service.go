package services

import (
	model "golang-hornet-task/models"
	"golang-hornet-task/repositories"
)

type DocumentService struct {
	Repo *repositories.DocumentRepository
}

func NewDocumentService(repo *repositories.DocumentRepository) *DocumentService {
	return &DocumentService{Repo: repo}
}

func (s *DocumentService) SaveDocument(doc model.Document) error {
	return s.Repo.Save(doc)
}

func (s *DocumentService) GetAllDocuments() ([]model.Document, error) {
	return s.Repo.GetAll()
}

func (s *DocumentService) GetDocumentByID(id int64) (model.Document, error) {
	return s.Repo.GetByID(id)
}

func (s *DocumentService) DeleteDocumentByID(id int64) bool {
	return s.Repo.DeleteById(id)
}

func (s *DocumentService) SearchDocuments(query string) ([]model.Document, error) {
	return s.Repo.SearchByNameLogic(query)
}

type DocumentServiceInterface interface {
	SaveDocument(doc model.Document) error
	GetAllDocuments() ([]model.Document, error)
	GetDocumentByID(id int64) (model.Document, error)
	DeleteDocumentByID(id int64) bool
	SearchDocuments(query string) ([]model.Document, error)
}
