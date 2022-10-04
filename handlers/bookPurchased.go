package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	bookpurchaseddto "waysbook/dto/bookPurchased"
	dto "waysbook/dto/result"
	"waysbook/models"
	"waysbook/repositories"
)

type handlerBookPurchased struct {
	BookPurchasedRepository repositories.BookPurchasedRepository
}

func HanlderBookPurchasedRepository(BookPurchasedRepository repositories.BookPurchasedRepository) *handlerBookPurchased {
	return &handlerBookPurchased{BookPurchasedRepository}
}

func (h *handlerBookPurchased) FindBooksPurcased(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	books, err := h.BookPurchasedRepository.FindBooksPurcased()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Error nya disini"}
		json.NewEncoder(w).Encode(response)
		return
	}

	filePath := os.Getenv("FILE_PATH")

	bookResponse := make([]bookpurchaseddto.BookPurchasedResponse, 0)
	for _, book := range books {
		bookResponse = append(bookResponse, bookpurchaseddto.BookPurchasedResponse{
			ID: book.ID,
			BookID: book.BookID,
			BookTitle: book.BookTitle,
			BookThumbnail: filePath + book.BookThumbnail,
		})
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: bookResponse}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBookPurchased) CreateBookPurchased(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := bookpurchaseddto.BookPurchasedRequest{
		BookID: r.FormValue("book_id"),
		BookTitle: r.FormValue("book_title"),
		BookThumbnail: r.FormValue("book_thumbnail"),
	}

	bookId, _ := strconv.Atoi(request.BookID)

	bookPurchased := models.BookPurchased{
		BookID: bookId,
		BookTitle: request.BookTitle,
		BookThumbnail: request.BookThumbnail,
	}

	data, err := h.BookPurchasedRepository.CreateBookPurchased(bookPurchased)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}