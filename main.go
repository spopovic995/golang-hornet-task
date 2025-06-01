package main

import (
	"golang-hornet-task/controllers"
	"golang-hornet-task/db"
	"golang-hornet-task/repositories"
	"golang-hornet-task/services"
	"log"
	"net/http"
)

func main() {

	// Initialize DB
	database := db.InitDB()

	// Create repository, service, and controller
	repo := repositories.NewDocumentRepository(database)
	service := services.NewDocumentService(repo)
	controller := controllers.NewDocumentController(service)

	// Register HTTP routes, here it checks for the method of the request and calls the appropriate function
	http.HandleFunc("/documents", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			controller.CreateDocument(w, r)
		case http.MethodGet:
			id := r.URL.Query().Get("id")
			if id != "" {
				controller.GetDocumentByID(w, r)
			} else {
				controller.GetAllDocuments(w, r)
			}
		case http.MethodDelete:
			controller.DeleteDocument(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/documents/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			controller.SearchDocuments(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
