package routes

import (
	"waysbook/handlers"
	psql "waysbook/pkg/dbConnection"
	"waysbook/pkg/middlewares"
	"waysbook/repositories"

	"github.com/gorilla/mux"
)

func BookRoutes(r *mux.Router) {
	bookRepository := repositories.RepositoryBook(psql.DB)
	h := handlers.HanlderBook(bookRepository)

	r.HandleFunc("/books", h.FindBooks).Methods("GET")
	r.HandleFunc("/book", middlewares.Auth(middlewares.IsSeller(middlewares.UploadPdf(h.CreateBook)))).Methods("POST")
	r.HandleFunc("/update-book-attachment/{id}", middlewares.Auth(middlewares.IsSeller(middlewares.UploadPdf(h.UpdateBookAttachment)))).Methods("POST")
	r.HandleFunc("/update-book-thumbnail/{id}", middlewares.Auth(middlewares.IsSeller(middlewares.UploadImage(h.UpdateBookThumbnail)))).Methods("POST")
}