package delivery

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type IHttpDelivery interface {
	MakeLoan(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	MakePayment(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	GetOutstanding(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	GetDeliquentUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	Repay(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
}
