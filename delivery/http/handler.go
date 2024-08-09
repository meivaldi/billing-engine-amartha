package http

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	delivery "github.com/meivaldi/billing-engine/delivery"
	"github.com/meivaldi/billing-engine/model"
	usecase "github.com/meivaldi/billing-engine/usecase"
)

type HttpHandler struct {
	userUsecase    usecase.IUserUsecase
	billingUsecase usecase.IBillingUsecase
	paymentUsecase usecase.IPaymentUsecase
}

const (
	LoanAmount = 5500000
)

func NewHttpHandler(
	userUc usecase.IUserUsecase,
	bilUC usecase.IBillingUsecase,
	payUC usecase.IPaymentUsecase,
) delivery.IHttpDelivery {
	return &HttpHandler{
		userUsecase:    userUc,
		billingUsecase: bilUC,
		paymentUsecase: payUC,
	}
}

func (h *HttpHandler) MakeLoan(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		userData        model.User
		nextPaymentDate time.Time
		responseData    HttpResponse
	)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responseData = HttpResponse{
			Err:     true,
			Message: err.Error(),
		}
		errorHandler(w, responseData)
		return
	}

	if err := json.Unmarshal(body, &userData); err != nil {
		responseData = HttpResponse{
			Err:     true,
			Message: err.Error(),
		}
		errorHandler(w, responseData)
		return
	}

	userID, err := h.userUsecase.CreateUser(userData)
	if err != nil {
		responseData = HttpResponse{
			Err:     true,
			Message: err.Error(),
		}
		errorHandler(w, responseData)
		return
	}

	userData.UserID = uint64(userID)
	nextPaymentDate = time.Now().Add(168 * time.Hour)

	billingData := model.Billing{
		UserID:            uint64(userID),
		OutstandingAmount: LoanAmount,
		NextPaymentDate:   nextPaymentDate, //for next 7 days
		IsDeliquent:       false,
	}

	id, err := h.billingUsecase.CreateLoan(billingData)
	if err != nil {
		responseData = HttpResponse{
			Err:     true,
			Message: err.Error(),
		}
		errorHandler(w, responseData)
		return
	}

	billingData.BillingID = uint64(id)
	responseData = HttpResponse{
		Err:     false,
		Message: "Successful create loan",
		Data: ResponseUserData{
			User:            userData,
			Amount:          LoanAmount,
			NextPaymentDate: nextPaymentDate,
		},
	}

	jsonResponse(w, responseData)
}

func (h *HttpHandler) MakePayment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		payment      model.Payment
		responseData HttpResponse
	)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responseData = HttpResponse{
			Err:     true,
			Message: err.Error(),
		}
		errorHandler(w, responseData)
		return
	}

	if err := json.Unmarshal(body, &payment); err != nil {
		responseData = HttpResponse{
			Err:     true,
			Message: err.Error(),
		}
		errorHandler(w, responseData)
		return
	}

	paymentId, err := h.paymentUsecase.MakePayment(payment)
	if err != nil {
		responseData = HttpResponse{
			Err:     true,
			Message: err.Error(),
		}
		errorHandler(w, responseData)
		return
	}

	payment.Id = uint64(paymentId)
	responseData = HttpResponse{
		Err:     false,
		Message: "Successful create payment",
		Data:    payment,
	}

	jsonResponse(w, responseData)
}

func (h *HttpHandler) GetOutstanding(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var responseData HttpResponse
	userId, err := strconv.Atoi(ps.ByName("userId"))
	if err != nil {
		responseData = HttpResponse{
			Err:     true,
			Message: err.Error(),
		}
		errorHandler(w, responseData)
		return
	}

	billing, err := h.billingUsecase.GetOutstanding(uint64(userId))
	if err != nil {
		responseData = HttpResponse{
			Err:     true,
			Message: err.Error(),
		}
		errorHandler(w, responseData)
		return
	}

	responseData = HttpResponse{
		Err:     false,
		Message: "successfully get outstanding data",
		Data:    billing,
	}
	jsonResponse(w, responseData)
}

func (h *HttpHandler) GetDeliquentUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var responseData HttpResponse

	users, err := h.userUsecase.GetDeliquentUsers()
	if err != nil {
		responseData = HttpResponse{
			Err:     true,
			Message: err.Error(),
		}
		errorHandler(w, responseData)
		return
	}

	responseData = HttpResponse{
		Err:     false,
		Message: "successfully get deliquent users",
		Data:    users,
	}

	jsonResponse(w, responseData)
}

func (h *HttpHandler) Repay(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		payment      model.Payment
		responseData HttpResponse
	)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responseData = HttpResponse{
			Err:     true,
			Message: err.Error(),
		}
		errorHandler(w, responseData)
		return
	}

	if err := json.Unmarshal(body, &payment); err != nil {
		responseData = HttpResponse{
			Err:     true,
			Message: err.Error(),
		}
		errorHandler(w, responseData)
		return
	}

	paymentIds, err := h.paymentUsecase.Repay(payment)
	if err != nil {
		responseData = HttpResponse{
			Err:     true,
			Message: err.Error(),
		}
		errorHandler(w, responseData)
		return
	}

	responseData = HttpResponse{
		Err:     false,
		Message: "Successful create payment",
		Data:    paymentIds,
	}

	jsonResponse(w, responseData)
}

func errorHandler(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	json.NewEncoder(w).Encode(data)
}

func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
