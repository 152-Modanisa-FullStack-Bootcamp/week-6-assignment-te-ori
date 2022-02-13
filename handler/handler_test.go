package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"my_account/account"
	"my_account/mock"
	"my_account/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	MIN_AMOUNT   = float32(-100)
	INIT_BALANCE = float32(0)
)

func TestPut(t *testing.T) {
	t.Run("put non-existing user", func(t *testing.T) {
		expectedBalance := &putResponse{Balance: INIT_BALANCE}
		username := "orhan"

		service := mock.NewMockIAccountService(gomock.NewController(t))

		service.EXPECT().
			Put(username).
			Return(INIT_BALANCE, nil).
			Times(1)

		handler := NewAccountHandler(service)

		r := httptest.NewRequest(http.MethodPut, "/orhan", nil)
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, r)
		var respondedBalance putResponse

		json.Unmarshal(w.Body.Bytes(), &respondedBalance)

		assert.Equal(t, expectedBalance, &respondedBalance)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.Equal(t, "application/json; charset=UTF-8", w.Header().Get("content-type"))
	})

	t.Run("put existing user", func(t *testing.T) {
		expectedBalance := &putResponse{Balance: float32(10)}
		username := "orhan"

		service := mock.NewMockIAccountService(gomock.NewController(t))

		service.EXPECT().
			Put(username).
			Return(float32(10), nil).
			Times(1)

		handler := NewAccountHandler(service)

		r := httptest.NewRequest(http.MethodPut, "/orhan", nil)
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, r)
		var respondedBalance putResponse

		json.Unmarshal(w.Body.Bytes(), &respondedBalance)

		assert.Equal(t, expectedBalance, &respondedBalance)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.Equal(t, "application/json; charset=UTF-8", w.Header().Get("content-type"))
	})
}
func TestAccountOf(t *testing.T) {
	t.Run("get existing user", func(t *testing.T) {
		expectedAccount := &account.Account{Username: "orhan", Balance: float32(10)}

		service := mock.NewMockIAccountService(gomock.NewController(t))

		service.EXPECT().
			AccountOf(expectedAccount.Username).
			Return(expectedAccount).
			Times(1)

		handler := NewAccountHandler(service)

		r := httptest.NewRequest(http.MethodGet, "/orhan", nil)
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, r)
		var respondedAccount account.Account

		json.Unmarshal(w.Body.Bytes(), &respondedAccount)

		assert.Equal(t, expectedAccount, &respondedAccount)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.Equal(t, "application/json; charset=UTF-8", w.Header().Get("content-type"))
	})

	t.Run("get nonexisting user", func(t *testing.T) {
		service := mock.NewMockIAccountService(gomock.NewController(t))

		service.EXPECT().
			AccountOf("orhan").
			Return(nil).
			Times(1)

		handler := NewAccountHandler(service)

		r := httptest.NewRequest(http.MethodGet, "/orhan", nil)
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, r)

		assert.Nil(t, w.Body.Bytes())
		assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
	})
}
func TestAccounts(t *testing.T) {
	t.Run("get existing users", func(t *testing.T) {
		expectedAccounts := []*account.Account{
			utils.CreateTestAccount("user1", 100),
			utils.CreateTestAccount("user2", -50),
			utils.CreateTestAccount("user3", 70),
			utils.CreateTestAccount("user4", 23.05),
		}

		service := mock.NewMockIAccountService(gomock.NewController(t))

		service.EXPECT().
			Accounts().
			Return(expectedAccounts).
			Times(1)

		handler := NewAccountHandler(service)

		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, r)
		var respondedAccounts []*account.Account

		json.Unmarshal(w.Body.Bytes(), &respondedAccounts)

		assert.Equal(t, expectedAccounts, respondedAccounts)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.Equal(t, "application/json; charset=UTF-8", w.Header().Get("content-type"))
	})
}

func TestUpdateBalance(t *testing.T) {
	t.Run("update nonexisting user", func(t *testing.T) {
		username := "orhan"
		service := mock.NewMockIAccountService(gomock.NewController(t))

		service.EXPECT().
			UpdateBalance(username, float32(10)).
			Return(float32(0), float32(0), fmt.Errorf("'%s' not found", username)).
			Times(1)

		handler := NewAccountHandler(service)

		json, _ := json.Marshal(updateBalanceRequest{Balance: float32(10)})

		r := httptest.NewRequest(http.MethodPost, "/orhan", bytes.NewReader(json))
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, r)

		assert.Equal(t, fmt.Sprintf("'%s' not found", username), w.Body.String())
		assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
	})

	t.Run("update existing user successfully", func(t *testing.T) {
		username := "orhan"
		service := mock.NewMockIAccountService(gomock.NewController(t))

		service.EXPECT().
			UpdateBalance(username, float32(10)).
			Return(float32(10), float32(110), nil).
			Times(1)

		handler := NewAccountHandler(service)

		requestBody, _ := json.Marshal(updateBalanceRequest{Balance: float32(10)})

		r := httptest.NewRequest(http.MethodPost, "/orhan", bytes.NewReader(requestBody))
		w := httptest.NewRecorder()

		var response updateBalanceResponse

		handler.ServeHTTP(w, r)
		err := json.Unmarshal(w.Body.Bytes(), &response)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.Equal(t, "", response.Message)
		assert.Equal(t, float32(10), response.NewBalance)
		assert.Equal(t, float32(110), response.Limit)
	})

	t.Run("update existing user exceed limit", func(t *testing.T) {
		username := "orhan"
		service := mock.NewMockIAccountService(gomock.NewController(t))

		service.EXPECT().
			UpdateBalance(username, float32(-10)).
			Return(float32(5), float32(5), fmt.Errorf("no sufficient balance")).
			Times(1)

		handler := NewAccountHandler(service)

		requestBody, _ := json.Marshal(updateBalanceRequest{Balance: float32(-10)})

		r := httptest.NewRequest(http.MethodPost, "/orhan", bytes.NewReader(requestBody))
		w := httptest.NewRecorder()

		var response updateBalanceResponse

		handler.ServeHTTP(w, r)
		err := json.Unmarshal(w.Body.Bytes(), &response)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
		assert.Equal(t, "no sufficient balance", response.Message)
		assert.Equal(t, float32(5), response.NewBalance)
		assert.Equal(t, float32(5), response.Limit)
	})
}

func TestUnsupportedMethods(t *testing.T) {

}
