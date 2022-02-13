package handler

import (
	"encoding/json"
	"fmt"
	"my_account/account"
	"net/http"
)

// type routeTable map[int]map[string]func(http.ResponseWriter, *http.Request)

type AccountHandler struct {
	service account.IAccountService
}

type putResponse struct {
	Balance float32 `json:"balance"`
}

type updateBalanceRequest struct {
	Balance float32 `json:"balance"`
}

type updateBalanceResponse struct {
	NewBalance float32 `json:"newBalance"`
	Limit      float32 `json:"limit"`
	Message    string  `json:"message"`
}

func (hnd *AccountHandler) put(w http.ResponseWriter, username string) {
	balance, err := hnd.service.Put(username)

	if err != nil {
		internalServerError(w, err)
		return
	}

	json, err := json.Marshal(putResponse{Balance: balance})
	if err != nil {
		internalServerError(w, err)
		return
	}

	w.Header().Add("content-type", "application/json; charset=UTF-8")
	w.Write(json)
}

func (hnd *AccountHandler) updateBalance(w http.ResponseWriter, username string, amount float32) {
	newBalance, limit, err := hnd.service.UpdateBalance(username, amount)

	response := &updateBalanceResponse{
		NewBalance: newBalance,
		Limit:      limit,
	}

	if err != nil {
		response.Message = err.Error()
		if response.Message == fmt.Sprintf("'%s' not found", username) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(response.Message))
		} else if response.Message == "no sufficient balance" {
			w.WriteHeader(http.StatusBadRequest)
			jsonResp, _ := json.Marshal(*response)
			jsonResponse(w, jsonResp)
		} else {
			internalServerError(w, err)
		}

		return
	}

	w.WriteHeader(http.StatusOK)
	jsonResp, _ := json.Marshal(*response)
	jsonResponse(w, jsonResp)
}

func (hnd *AccountHandler) accountOf(w http.ResponseWriter, username string) {
	user := hnd.service.AccountOf(username)

	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json, err := json.Marshal(user)
	if err != nil {
		internalServerError(w, err)
		return
	}

	jsonResponse(w, json)
}

func (hnd *AccountHandler) accounts(w http.ResponseWriter) {
	accounts := hnd.service.Accounts()

	json, err := json.Marshal(accounts)
	if err != nil {
		internalServerError(w, err)
		return
	}

	w.Header().Add("content-type", "application/json; charset=UTF-8")
	w.Write(json)
}

func (hnd *AccountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" && r.Method == http.MethodGet {
		hnd.accounts(w)
	} else if r.URL.Path != "" {
		username := r.URL.Path[1:]
		if r.Method == http.MethodGet {
			hnd.accountOf(w, username)
		} else if r.Method == http.MethodPut {
			hnd.put(w, username)
		} else if r.Method == http.MethodPost {
			var req updateBalanceRequest
			dec := json.NewDecoder(r.Body)

			err := dec.Decode(&req)

			if err != nil {
				internalServerError(w, err)
				return
			}

			hnd.updateBalance(w, username, req.Balance)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func NewAccountHandler(service account.IAccountService) *AccountHandler {
	return &AccountHandler{service: service}
}

func internalServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func jsonResponse(w http.ResponseWriter, resp []byte) {
	w.Header().Add("content-type", "application/json; charset=UTF-8")
	w.Write(resp)
}
