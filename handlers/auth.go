package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
	authdto "waysbook/dto/auth"
	dto "waysbook/dto/result"
	"waysbook/models"
	"waysbook/pkg/bcrypt"
	jwtToken "waysbook/pkg/jwt"
	"waysbook/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get Request Body
	request := new(authdto.RegisterRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Validate Request
	validation := validator.New()
	if err := validation.Struct(request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Check if email already exist
	err = h.AuthRepository.CheckEmailExist(request.Email)

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Email already exist"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create New User
	user := models.User{
		Email: request.Email,
		FullName: request.FullName,
		Password: password,
		Gender: "-",
		Address: "-",
		IsSeller: false,
		Phone: "-",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	data, err := h.AuthRepository.Register(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Generate Token
	claims := jwt.MapClaims{}
	claims["id"] = data.ID
	claims["is_seller"] = data.IsSeller
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 hour expired

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(err)
		fmt.Println("Unauthorized")
		return
	}

	registerResponse := authdto.AuthResponse {
		ID: data.ID,
		FullName: data.FullName,
		Email: data.Email,
		IsSeller: data.IsSeller,
		Gender: data.Gender,
		Phone: data.Phone,
		Address: data.Address,
		Token: token,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: registerResponse}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerAuth) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(authdto.LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user := models.User{
		Email: request.Email,
		Password: request.Password,
	}

	data, err := h.AuthRepository.Login(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Email is incorrect"}
		json.NewEncoder(w).Encode(response)
		return
	}

	isValid := bcrypt.CheckPasswordHash(request.Password, data.Password)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Password is incorrect"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Generate Token
	claims := jwt.MapClaims{}
	claims["id"] = data.ID
	claims["is_seller"] = data.IsSeller
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 hour expired

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if err != nil {
		log.Println(err)
		fmt.Println("Unauthorized")
		return
	}

	if errGenerateToken != nil {
		log.Println(err)
		fmt.Println("Unauthorized")
		return
	}

	loginResponse := authdto.AuthResponse{
		ID: data.ID,
		FullName: data.FullName,
		Email: data.Email,
		IsSeller: data.IsSeller,
		Gender: data.Gender,
		Phone: data.Phone,
		Address: data.Address,
		Token: token,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: loginResponse}
	json.NewEncoder(w).Encode(response)
}