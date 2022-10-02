package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	cartdto "waysbook/dto/cart"
	dto "waysbook/dto/result"
	"waysbook/models"
	"waysbook/repositories"

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

	// filePath := os.Getenv("FILE_PATH")

	cartResponse := make([]cartdto.CartResponse, 0)
	for _, cart := range carts {
		cartResponse = append(cartResponse, cartdto.CartResponse{
			ID: cart.ID,
			Price: cart.Price,
			BookID: cart.BookID,
			UserID: cart.UserID,
			SellerID: cart.SellerID,
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

	request := cartdto.CartRequest{
		Price: r.FormValue("price"),
		BookID: r.FormValue("book_id"),
		SellerID: r.FormValue("seller_id"),
	}

	price, _ := strconv.Atoi(request.Price)
	bookId, _ := strconv.Atoi(request.BookID)
	sellerId, _ := strconv.Atoi(request.SellerID)

	cart := models.Cart{
		UserID: userId,
		Price: price,
		BookID: bookId,
		SellerID: sellerId,
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

	cartResponse := cartdto.CartResponse{
		ID: newCart.ID,
		Price: newCart.Price,
		SellerID: newCart.SellerID,
		BookID: newCart.BookID,
		UserID: newCart.UserID,
	}

	w.WriteHeader(http.StatusCreated)
	response := dto.SuccessResult{Code: http.StatusCreated, Data: cartResponse}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCart) DeleteCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cartId, _ := strconv.Atoi(r.FormValue("cart_id"))
	cart, err := h.CartRepository.GetCartByID(cartId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
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

	cartResponse := cartdto.CartResponse{
		ID: data.ID,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: cartResponse}
	json.NewEncoder(w).Encode(response)
}