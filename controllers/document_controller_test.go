package controllers

import (
	"bytes"
	"encoding/json"
	"golang-hornet-task/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockDocumentService struct {
	SaveDocumentFn func(doc models.Document) error
}

func (m *MockDocumentService) SaveDocument(doc models.Document) error {
	if m.SaveDocumentFn != nil {
		return m.SaveDocumentFn(doc)
	}
	return nil
}

func (m *MockDocumentService) GetAllDocuments() ([]models.Document, error) {
	return nil, nil
}

func (m *MockDocumentService) GetDocumentByID(id int64) (models.Document, error) {
	return models.Document{}, nil
}

func (m *MockDocumentService) DeleteDocumentByID(id int64) bool {
	return true
}

func (m *MockDocumentService) SearchDocuments(query string) ([]models.Document, error) {
	return nil, nil
}

func TestCreateDocument(t *testing.T) {
	mockService := &MockDocumentService{
		SaveDocumentFn: func(doc models.Document) error {
			assert.Equal(t, "test name", doc.Name)
			return nil
		},
	}
	controller := NewDocumentController(mockService)

	doc := models.Document{
		Name: "test name",
	}

	body, err := json.Marshal(doc)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/documents", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	controller.CreateDocument(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, "Document created successfully", rr.Body.String())
}
