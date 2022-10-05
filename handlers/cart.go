package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"
	cartdto "waysbook/dto/cart"
	dto "waysbook/dto/result"
	"waysbook/models"
	"waysbook/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerCart struct {
	CartRepository repositories.CartRepository
}

func HandlerCart(CartRepository repositories.CartRepository) *handlerCart {
	return &handlerCart{CartRepository}
}

func (h *handlerCart) FindCartsByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get User ID from token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	carts, err := h.CartRepository.FindCartsByUserID(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	filePath := os.Getenv("FILE_PATH")

	cartResponse := make([]cartdto.CartResponse, 0)
	for _, cart := range carts {
		cartResponse = append(cartResponse, cartdto.CartResponse{
			ID: cart.ID,
			QTY: cart.QTY,
			Subtotal: cart.Subtotal,
			SellerID: cart.SellerID,
			BookID: cart.BookID,
			UserID: userId,
			BookTitle: cart.Book.Title,
			Slug: cart.Book.Slug,
			BookThumbnail: filePath + cart.Book.Thumbnail,
			Author: cart.Book.Author,
		})
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: cartResponse}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCart) AddCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get user from token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	request := cartdto.CreateCartRequest{
		SellerID: r.FormValue("seller_id"),
		BookID: r.FormValue("book_id"),
		SubTotal: r.FormValue("subtotal"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	sellerId, _ := strconv.Atoi(request.SellerID)
	bookId, _ := strconv.Atoi(request.BookID)
	subtotal, _ := strconv.Atoi(request.SubTotal)

	cart := models.Cart{
		QTY: 1,
		Subtotal: subtotal,
		BookID: bookId,
		SellerID: sellerId,
		UserID: userId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	newCart, err := h.CartRepository.AddCart(cart)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := dto.SuccessResult{Code: http.StatusCreated, Data: newCart}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCart) DeleteCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := cartdto.DeleteCartRequest{
		ID: r.FormValue("cart_id"),
		UserID: r.FormValue("user_id"),
	}

	cartId, _ := strconv.Atoi(request.ID)
	userId, _ := strconv.Atoi(request.UserID)

	cart, err := h.CartRepository.GetCartByID(int(cartId))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Error nya disini"}
		json.NewEncoder(w).Encode(response)
		return
	}

	if cart.UserID != int(userId) {
		w.WriteHeader(http.StatusForbidden)
		response := dto.ErrorResult{Code: http.StatusForbidden, Message: "You can't delete someone else's cart! Please just delete yours"}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.CartRepository.DeleteCart(cart)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// allCarts, err := h.CartRepository.FindCartsByUserID()

	cartResponse := cartdto.DeleteCartResponse{
		ID: data.ID,
		Message: "Deleting cart success!",
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: cartResponse}
	json.NewEncoder(w).Encode(response)
}