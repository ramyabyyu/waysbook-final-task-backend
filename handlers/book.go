package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"
	bookdto "waysbook/dto/book"
	dto "waysbook/dto/result"
	"waysbook/models"
	dateformat "waysbook/pkg/dateFormat"
	"waysbook/pkg/slug"
	"waysbook/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerBook struct {
	BookRepository repositories.BookRepository
}

func HanlderBook(BookRepository repositories.BookRepository) *handlerBook {
	return &handlerBook{BookRepository}
}

func (h *handlerBook) FindBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	books, err := h.BookRepository.FindBooks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	filePath := os.Getenv("FILE_PATH")

	bookResponse := make([]bookdto.BookResponse, 0)
	for _, book := range books {
		bookResponse = append(bookResponse, bookdto.BookResponse{
			ID: book.ID,
			Title: book.Title,
			Slug: book.Slug,
			PublicationDate: book.PublicationDate,
			Pages: book.Pages,
			ISBN: book.ISBN,
			Price: book.Price,
			IsPromo: book.IsPromo,
			Discount: book.Discount,
			Description: book.Description,
			BookAttachment: filePath + book.BookAttachment,
			Thumbnail: filePath + book.Thumbnail,
			UserID: book.UserID,
		})
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: bookResponse}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBook) CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get file name
	dataContext := r.Context().Value("dataPdf")
	filename := dataContext.(string)

	request := bookdto.CreateBookRequest{
		Title: r.FormValue("title"),
		PublicationDate: r.FormValue("publication_date"),
		Pages: r.FormValue("pages"),
		ISBN: r.FormValue("ISBN"),
		Price: r.FormValue("price"),
		Description: r.FormValue("description"),
		BookAttachment: r.FormValue("document"),
	}

	title := request.Title

	publicationDate, _ := dateformat.ConvertStrToDate(request.PublicationDate)
	pages, _ := strconv.Atoi(request.Pages)
	isbn, _ := strconv.Atoi(request.ISBN)
	price, _ := strconv.Atoi(request.Price)

	// Get user id from token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	book := models.Book{
		Title: title,
		Thumbnail: "-",
		Slug: slug.GenerateSlug(title),
		Description: request.Description,
		PublicationDate: publicationDate,
		Pages: int(pages),
		ISBN: int(isbn),
		Price: int(price),
		IsPromo: false,
		Discount: 0,
		PriceAfterDiscount: 0,
		UserID: userId,
		BookAttachment: filename,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	newBook, err := h.BookRepository.CreateBook(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	filePath := os.Getenv("FILE_PATH")

	bookResponse := bookdto.BookResponse{
		ID: newBook.ID,
		Title: newBook.Title,
		Slug: newBook.Slug,
		PublicationDate: newBook.PublicationDate,
		Pages: newBook.Pages,
		ISBN: newBook.ISBN,
		Price: newBook.Price,
		Description: newBook.Description,
		BookAttachment: filePath + newBook.BookAttachment,
		Thumbnail: filePath + newBook.Thumbnail,
		UserID: newBook.UserID,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: bookResponse}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBook) UpdateBookThumbnail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get ID
	bookId, _ := strconv.Atoi(mux.Vars(r)["id"])

	// Get user id from token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	bookAuthorId, err := h.BookRepository.GetBookByID(bookId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Id not found"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// check if bookAuthorId and userId from token is match
	if bookAuthorId.UserID == userId {
		w.WriteHeader(http.StatusForbidden)
		response := dto.ErrorResult{Code: http.StatusForbidden, Message: "You are not allowed to update this book"}
		json.NewEncoder(w).Encode(response)
	}
	// Get file name
	dataContext := r.Context().Value("dataImage")
	filename := dataContext.(string)

	request := bookdto.UpdateBookThumbnailRequest{
		Thumbnail: filename,
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	updatedBook, err := h.BookRepository.UpdateBookThumbnail(bookId, filename)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	filePath := os.Getenv("FILE_PATH")

	bookResponse := bookdto.UpdateBookThumbnailResponse{
		Thumbnail: filePath + updatedBook.Thumbnail,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: bookResponse}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBook) GetUserBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	books, err := h.BookRepository.GetUserBook(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "ID not found"}
		json.NewEncoder(w).Encode(response)
		return
	}

	filePath := os.Getenv("FILE_PATH")

	bookResponse := make([]bookdto.BookResponse, 0)
	for _, book := range books {
		bookResponse = append(bookResponse, bookdto.BookResponse{
			ID: book.ID,
			Title: book.Title,
			Slug: book.Slug,
			PublicationDate: book.PublicationDate,
			Pages: book.Pages,
			ISBN: book.ISBN,
			Price: book.Price,
			IsPromo: book.IsPromo,
			Discount: book.Discount,
			PriceAfterDiscount: book.PriceAfterDiscount,
			Description: book.Description,
			BookAttachment: filePath + book.BookAttachment,
			Thumbnail: filePath + book.Thumbnail,
			UserID: book.UserID,
		})
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: bookResponse}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBook) UpdateBookPromo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := bookdto.UpdateBookPromoRequest{
		ID: r.FormValue("book_id"),
		Discount: r.FormValue("discount"),
	}

	bookId, _ := strconv.Atoi(request.ID)
	discount, _ := strconv.Atoi(request.Discount)

	book, err := h.BookRepository.UpdateBookPromo(bookId, discount)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	filePath := os.Getenv("FILE_PATH")

	bookResponse := bookdto.BookResponse{
		ID: book.ID,
		Title: book.Title,
		Slug: book.Slug,
		PublicationDate: book.PublicationDate,
		Pages: book.Pages,
		ISBN: book.ISBN,
		Price: book.Price,
		IsPromo: book.IsPromo,
		Discount: book.Discount,
		PriceAfterDiscount: book.PriceAfterDiscount,
		Description: book.Description,
		BookAttachment: filePath + book.BookAttachment,
		Thumbnail: filePath + book.Thumbnail,
		UserID: book.UserID,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: bookResponse}
	json.NewEncoder(w).Encode(response)
}

