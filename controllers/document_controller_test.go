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

// A mock of the service, it's a placeholder for the real service
type MockDocumentService struct {
	// This i understand as a placeholder for a function, not yet defined
	SaveDocumentFn func(doc models.Document) error
}

// This type of func declaration means that the MockDocumentService is like injected into this method (Java explanation)
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
		//Here you define the funcion's implementation
		SaveDocumentFn: func(doc models.Document) error {
			assert.Equal(t, "test name", doc.Name)
			assert.Equal(t, "test description", doc.Description)
			return nil
		},
	}
	controller := NewDocumentController(mockService)

	doc := models.Document{
		Name:        "test name",
		Description: "test description",
	}

	body, err := json.Marshal(doc)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/documents", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	// Here you call the CreateDocument which calls the SaveDocument function that we mocked earlier, not the real one
	controller.CreateDocument(rr, req)
	//The following code is executed exactly as it is in the controller
	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, "Document created successfully", rr.Body.String())
}
