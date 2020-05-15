package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/saefullohmaslul/golang-example/src/apps"
	db "github.com/saefullohmaslul/golang-example/src/database"
	"github.com/saefullohmaslul/golang-example/src/middlewares/exception"
	"github.com/saefullohmaslul/golang-example/src/repositories"
	"github.com/saefullohmaslul/golang-example/src/utils"
	"github.com/saefullohmaslul/golang-example/src/utils/flag"
	"github.com/stretchr/testify/assert"
)

type createUserSuccess struct {
	utils.Response
	Result repositories.GetUser `json:"result"`
}

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

		actual := createUserSuccess{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, flag.CreateUserSuccess.Message, actual.Message)
		assert.Equal(t, http.StatusOK, actual.Status)
		assert.NotEmpty(t, actual.Result)
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

		actual := exception.Exception{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, http.StatusBadRequest, actual.Status)
		assert.Equal(t, actual.Flag, flag.CreateUserAlreadyExist.Flag)
		assert.NotEmpty(t, actual.Errors)
		assert.Equal(t, flag.CreateUserAlreadyExist.Error.Message, actual.Errors.Message)
		assert.Equal(t, flag.CreateUserAlreadyExist.Error.Flag, actual.Errors.Flag)
	})

	t.Run("it should return invalid body with invalid name format", func(t *testing.T) {
		defer db.DropAllTable()
		body := map[string]interface{}{
			"name":     123,
			"email":    "user@email.com",
			"password": "123456",
		}
		w, _ := initTestCreateUser(body)

		actual := exception.Exception{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, http.StatusBadRequest, actual.Status)
		assert.Equal(t, actual.Flag, flag.CreateUserInvalidBody.Flag)
		assert.NotEmpty(t, actual.Errors)
		assert.NotEmpty(t, actual.Errors.Message)
		assert.Equal(t, flag.CreateUserInvalidBody.Error.Flag, actual.Errors.Flag)
	})

	t.Run("it should return invalid body with invalid email format", func(t *testing.T) {
		defer db.DropAllTable()
		body := map[string]interface{}{
			"name":     "User Test",
			"email":    "useremail.com",
			"password": "123456",
		}
		w, _ := initTestCreateUser(body)

		actual := exception.Exception{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, http.StatusBadRequest, actual.Status)
		assert.Equal(t, actual.Flag, flag.CreateUserInvalidBody.Flag)
		assert.NotEmpty(t, actual.Errors)
		assert.NotEmpty(t, actual.Errors.Message)
		assert.Equal(t, flag.CreateUserInvalidBody.Error.Flag, actual.Errors.Flag)
	})

	t.Run("it should return invalid body with invalid password format", func(t *testing.T) {
		defer db.DropAllTable()
		body := map[string]interface{}{
			"name":  "User Test",
			"email": "user@email.com",
		}
		w, _ := initTestCreateUser(body)

		actual := exception.Exception{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, http.StatusBadRequest, actual.Status)
		assert.Equal(t, actual.Flag, flag.CreateUserInvalidBody.Flag)
		assert.NotEmpty(t, actual.Errors)
		assert.NotEmpty(t, actual.Errors.Message)
		assert.Equal(t, flag.CreateUserInvalidBody.Error.Flag, actual.Errors.Flag)
	})
}
