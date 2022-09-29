package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	dto "waysbook/dto/result"
	userdto "waysbook/dto/user"
	"waysbook/repositories"

	"github.com/golang-jwt/jwt/v4"
)

type handlerUser struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handlerUser {
	return &handlerUser{UserRepository}
}

func (h *handlerUser) GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := h.UserRepository.GetAllUser()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	filePath := os.Getenv("FILE_PATH")

	userResponse := make([]userdto.UserResponse, 0)
	for _, user := range users {
		userResponse = append(userResponse, userdto.UserResponse{
			ID: user.ID,
			FullName: user.FullName,
			Email: user.Email,
			IsSeller: user.IsSeller,
			Gender: user.Gender,
			Phone: user.Phone,
			Photo: filePath + user.Photo,
			IsPhotoChange: user.IsPhotoChange,
		})
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: userResponse}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get user id from token
	// Get user id from token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	user, err := h.UserRepository.GetUserByID(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	filePath := os.Getenv("FILE_PATH")

	userResponse := userdto.UserResponse{
		ID: user.ID,
		FullName: user.FullName,
		Email: user.Email,
		IsSeller: user.IsSeller,
		Gender: user.Gender,
		Phone: user.Phone,
		Photo: filePath + user.Photo,
		Address: user.Address,
		IsPhotoChange: user.IsPhotoChange,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: userResponse}
	json.NewEncoder(w).Encode(response)
}