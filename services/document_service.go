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
