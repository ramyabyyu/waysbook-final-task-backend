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
	r.HandleFunc("/book/{slug}", middlewares.Auth(h.GetBookBySlug)).Methods("GET")
	r.HandleFunc("/update-book-thumbnail/{id}", middlewares.Auth(middlewares.IsSeller(middlewares.UploadImage(h.UpdateBookThumbnail)))).Methods("POST")
	r.HandleFunc("/get-user-book", middlewares.Auth(middlewares.IsSeller(h.GetUserBook))).Methods("GET")
	r.HandleFunc("/update-book-promo", middlewares.Auth(middlewares.IsSeller(h.UpdateBookPromo))).Methods("POST")
	r.HandleFunc("/get-books-promo", h.GetBooksByPromo).Methods("GET")
}