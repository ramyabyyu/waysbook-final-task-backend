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

	"github.com/golang-jwt/jwt/v4"
)

type handlerBookPurchased struct {
	BookPurchasedRepository repositories.BookPurchasedRepository
}

func HanlderBookPurchased(BookPurchasedRepository repositories.BookPurchasedRepository) *handlerBookPurchased {
	return &handlerBookPurchased{BookPurchasedRepository}
}

func (h *handlerBookPurchased) FindBooksPurcased(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get user from token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))
	
	books, err := h.BookPurchasedRepository.FindBooksPurcased(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	filePath := os.Getenv("FILE_PATH")

	bookPurchasedResponse := make([]bookpurchaseddto.BookPurchasedResponse, 0)
	for _, book := range books {
		bookDetail, _ := h.BookPurchasedRepository.GetOneBook(book.ID)
		bookPurchasedResponse = append(bookPurchasedResponse, bookpurchaseddto.BookPurchasedResponse{
			ID: book.ID,
			UserID: userId,
			BookID: book.ID,
			BookTitle: bookDetail.Title,
			BookThumbnail: filePath + bookDetail.Thumbnail,
		})
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: bookPurchasedResponse}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBookPurchased) CreateBookPurchased(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userId, _ := strconv.Atoi(r.FormValue("user_id"))
	book_id1, _ := strconv.Atoi(r.FormValue("book_id1"))
	book_id2, _ := strconv.Atoi(r.FormValue("book_id2"))
	book_id3, _ := strconv.Atoi(r.FormValue("book_id3"))

	books := make([]models.BookPurchased, 0)
	books = append(books, models.BookPurchased{ UserID: userId, BookID: book_id1 })
	books = append(books, models.BookPurchased{ UserID: userId, BookID: book_id2 })
	books = append(books, models.BookPurchased{ UserID: userId, BookID: book_id3 })

	for i, _ := range books {
		h.BookPurchasedRepository.CreateBookPurchased(books[i])
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessNoDataResult{Code: http.StatusOK, Message: "Book Purchased Added Successfully"}
	json.NewEncoder(w).Encode(response)
}