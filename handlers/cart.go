package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	cartdto "waysbook/dto/cart"
	dto "waysbook/dto/result"
	"waysbook/models"
	"waysbook/repositories"

	"github.com/golang-jwt/jwt/v4"
)

type handlerCart struct {
	CartRepository repositories.CartRepository
}

func HanlderCart(CartRepository repositories.CartRepository) *handlerCart {
	return &handlerCart{CartRepository}
}

func (h *handlerCart) FindCartItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := cartdto.GetCartItemRequest{
		CartID: r.FormValue("cart_id"),
	}

	cartId, _ := strconv.Atoi(request.CartID)

	cartItems, err := h.CartRepository.FindCartItems(cartId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Cart not found"}
		json.NewEncoder(w).Encode(response)
		return
	}

	cartResponse := make([]cartdto.CartItemResponse, 0)
	for _, cartItem := range cartItems {
		cartResponse = append(cartResponse, cartdto.CartItemResponse{
			ID: cartItem.ID,
			BookID: cartItem.BookID,
			BookTitle: cartItem.BookTitle,
			BookThumbnail: cartItem.BookThumbnail,
			Price: cartItem.Price,
			Qty: cartItem.Qty,
			UserID: cartItem.Cart.UserID,
			BuyerName: cartItem.BuyerName,
			SellerID: cartItem.SellerID,
			SellerName: cartItem.SellerName,
		})
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: cartResponse}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCart) CreateCartItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// Get BookID from request
	request := cartdto.CreateCartItemRequest{
		BookID: r.FormValue("book_id"),
		CartID: r.FormValue("cart_id"),
	}

	// convert bookID and cartID from string to int
	bookId, _ := strconv.Atoi(request.BookID)
	cartId, _ := strconv.Atoi(request.CartID)

	// Get Book Data
	book, err := h.CartRepository.GetBook(bookId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Book not found"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Get User Data
	user, errUser := h.CartRepository.GetUser(userId)
	if errUser != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	filePath := os.Getenv("FILE_PATH")

	// check if book has promo
	var bookPrice int
	if book.PriceAfterDiscount == 0 {
		bookPrice = book.Price
	} else {
		bookPrice = book.PriceAfterDiscount
	}

	newCartItem := models.CartItem{
		CartID: cartId,
		BookID: book.ID,
		BookTitle: book.Title,
		BookThumbnail: filePath + book.Thumbnail,
		UserID: user.ID,
		BuyerName: user.FullName,
		SellerID: book.UserID,
		SellerName: book.User.FullName,
		Qty: 1,
		Price: bookPrice,
	}

	data, err := h.CartRepository.CreateCartItem(newCartItem)
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