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

	// Register HTTP route
	http.HandleFunc("/documents", controller.CreateDocument)

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
