package controllers

import (
	"encoding/json"
	"golang-hornet-task/models"
	"golang-hornet-task/services"
	"net/http"
)

// DocumentController depends on the interface, so it can accept mocks in tests
type DocumentController struct {
	Service services.DocumentServiceInterface
}

func NewDocumentController(service services.DocumentServiceInterface) *DocumentController {
	return &DocumentController{Service: service}
}

func (c *DocumentController) CreateDocument(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var doc models.Document
	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = c.Service.SaveDocument(doc)
	if err != nil {
		http.Error(w, "Failed to save document", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Document created successfully"))
}

// Other methods omitted for brevity as you only requested CreateDocument test
