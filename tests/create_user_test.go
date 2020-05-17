package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/saefullohmaslul/golang-example/src/utils/response"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"

	"github.com/saefullohmaslul/golang-example/src/apps"
	db "github.com/saefullohmaslul/golang-example/src/database"
	"github.com/saefullohmaslul/golang-example/src/utils/flag"
)

func initTestCreateUser(body map[string]interface{}) (*httptest.ResponseRecorder, *gin.Engine) {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	w := httptest.NewRecorder()
	b, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, "/user", strings.NewReader(string(b)))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)
	return w, r
}

func TestCreateUser(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer db.DropAllTable()
		body := map[string]interface{}{
			"name":     "User Test",
			"email":    "user@email.com",
			"password": "123456",
		}
		w, _ := initTestCreateUser(body)

		actual := response.Response{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, flag.CreateUserSuccess.Message, actual.Message)
		assert.Equal(t, http.StatusOK, actual.Status)
		assert.NotEmpty(t, actual.Data)
		assert.Empty(t, actual.Errors)
	})

	t.Run("it should return user already exist", func(t *testing.T) {
		defer db.DropAllTable()
		body := map[string]interface{}{
			"name":     "User Test",
			"email":    "user@email.com",
			"password": "123456",
		}
		_, r := initTestCreateUser(body)

		w := httptest.NewRecorder()
		b, _ := json.Marshal(body)

		req, _ := http.NewRequest(http.MethodPost, "/user", strings.NewReader(string(b)))
		req.Header.Set("Content-Type", "application/json")

		r.ServeHTTP(w, req)

		actual := response.Response{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusConflict, w.Code)
		assert.Equal(t, http.StatusConflict, actual.Status)
		assert.Equal(t, "User conflict", actual.Message)
		assert.Equal(t, "User with this email already exist", actual.Errors[0].Message)
		assert.Equal(t, "USER_ALREADY_EXIST", actual.Errors[0].Flag)
	})

	t.Run("it should return invalid body with invalid name format", func(t *testing.T) {
		defer db.DropAllTable()
		body := map[string]interface{}{
			"name":     123,
			"email":    "user@email.com",
			"password": "123456",
		}
		w, _ := initTestCreateUser(body)

		actual := response.Response{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, http.StatusBadRequest, actual.Status)
		assert.Equal(t, "Validation error", actual.Message)
		assert.Equal(t, "INVALID_BODY", actual.Errors[0].Flag)
		assert.NotEmpty(t, actual.Errors[0].Message)
	})

	t.Run("it should return invalid body with invalid email format", func(t *testing.T) {
		defer db.DropAllTable()
		body := map[string]interface{}{
			"name":     "User Test",
			"email":    "useremail.com",
			"password": "123456",
		}
		w, _ := initTestCreateUser(body)

		actual := response.Response{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, http.StatusBadRequest, actual.Status)
		assert.Equal(t, "Validation error", actual.Message)
		assert.Equal(t, "INVALID_BODY", actual.Errors[0].Flag)
		assert.NotEmpty(t, actual.Errors[0].Message)
	})

	t.Run("it should return invalid body with invalid password format", func(t *testing.T) {
		defer db.DropAllTable()
		body := map[string]interface{}{
			"name":  "User Test",
			"email": "user@email.com",
		}
		w, _ := initTestCreateUser(body)

		actual := response.Response{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, http.StatusBadRequest, actual.Status)
		assert.Equal(t, "Validation error", actual.Message)
		assert.Equal(t, "INVALID_BODY", actual.Errors[0].Flag)
		assert.NotEmpty(t, actual.Errors[0].Message)
	})
}
