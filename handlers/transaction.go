package handlers

// import (
// 	"encoding/json"
// 	"net/http"
// 	"os"
// 	"strconv"
// 	"time"
// 	dto "waysbook/dto/result"
// 	transactiondto "waysbook/dto/transaction"
// 	"waysbook/models"
// 	"waysbook/repositories"

// 	"github.com/golang-jwt/jwt/v4"
// 	"github.com/midtrans/midtrans-go"
// 	"github.com/midtrans/midtrans-go/coreapi"
// 	"github.com/midtrans/midtrans-go/snap"
// )

// var c = coreapi.Client{
// 	ServerKey: os.Getenv("SERVER_KEY"),
// 	ClientKey: os.Getenv("CLIENT_KEY"),
// }

// type handlerTransaction struct {
// 	TransactionRepository repositories.TransactionRepository
// }

// func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
// 	return &handlerTransaction{TransactionRepository}
// }

// func (h *handlerTransaction) CreateTransaction(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
// 	userId := int(userInfo["id"].(float64))

// 	request := transactiondto.TransactionRequest{
// 		CartID: r.FormValue("cart_id"),
// 		Total: r.FormValue("total"),
// 	}

// 	var TransIdIsMatch = false
// 	var TransactionId int
// 	for !TransIdIsMatch {
// 		TransactionId = int(time.Now().Unix()) * 2
// 		transactionData, _ := h.TransactionRepository.GetTransaction(TransactionId)
// 		if transactionData.ID == 0 {
// 			TransIdIsMatch = true
// 		}
// 	}

// 	cartId, _ := strconv.Atoi(request.CartID)
// 	total, _ := strconv.Atoi(request.Total)

// 	newTransaction := models.Transaction{
// 		ID: TransactionId,
// 		Total: total,
// 		CartID: cartId,
// 		Status: "success",
// 		UserID: userId,
// 	}

// 	data, err := h.TransactionRepository.CreateTransaction(newTransaction)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	dataTransaction, err := h.TransactionRepository.GetTransaction(data.ID)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	var s = snap.Client{}
// 	s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)

// 	req := &snap.Request{
// 		TransactionDetails: midtrans.TransactionDetails{
// 			OrderID: strconv.Itoa(dataTransaction.ID),
// 			GrossAmt: int64(dataTransaction.Total),
// 		},
// 		CreditCard: &snap.CreditCardDetails{
// 			Secure: true,
// 		},
// 		CustomerDetail: &midtrans.CustomerDetails{
// 			FName: dataTransaction.User.FullName,
// 			Email: dataTransaction.User.Email,
// 		},
// 	}

// 	snapResp, _ := s.CreateTransaction(req)

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Code: http.StatusOK, Data: snapResp}
// 	json.NewEncoder(w).Encode(response)
// }