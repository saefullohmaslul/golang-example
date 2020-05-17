package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"

	"github.com/saefullohmaslul/golang-example/src/apps"
	db "github.com/saefullohmaslul/golang-example/src/database"
	"github.com/saefullohmaslul/golang-example/src/database/entity"
	"github.com/saefullohmaslul/golang-example/src/repositories"
	"github.com/saefullohmaslul/golang-example/src/utils/flag"
	"github.com/saefullohmaslul/golang-example/src/utils/response"
)

func initTestUpdateUser(id string, body map[string]interface{}) (*httptest.ResponseRecorder, *gin.Engine) {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	w := httptest.NewRecorder()
	b, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPatch, "/user/"+id, strings.NewReader(string(b)))
	req.Header.Set("Content-Type", "application/json")

	userRepository := repositories.UserRepository{Conn: db.GetDB().Table("users")}
	userRepository.CreateUser(entity.User{
		ID:       1,
		Address:  "Jakarta",
		Age:      20,
		Email:    "email@email.com",
		Name:     "Saefulloh Maslul",
		Password: "123456",
	})

	r.ServeHTTP(w, req)
	return w, r
}

func TestUpdateUser(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer db.DropAllTable()
		body := map[string]interface{}{
			"age": 35,
		}
		w, _ := initTestUpdateUser("1", body)

		actual := response.Response{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, flag.UpdateUserSuccess.Message, actual.Message)
		assert.Equal(t, http.StatusOK, actual.Status)
		assert.NotEmpty(t, actual.Data)
		assert.Empty(t, actual.Errors)
	})

	t.Run("it should return user not exist", func(t *testing.T) {
		defer db.DropAllTable()
		body := map[string]interface{}{
			"age": 35,
		}
		w, _ := initTestUpdateUser("2", body)

		actual := response.Response{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, http.StatusNotFound, actual.Status)
		assert.Equal(t, "User not exist", actual.Message)
		assert.Equal(t, "USER_NOT_FOUND", actual.Errors[0].Flag)
		assert.Equal(t, "User with this ID not found", actual.Errors[0].Message)
	})

	t.Run("it should return invalid body with invalid email format", func(t *testing.T) {
		defer db.DropAllTable()
		body := map[string]interface{}{
			"email": "bukanatemaildotcom",
		}
		w, _ := initTestUpdateUser("1", body)

		actual := response.Response{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, http.StatusBadRequest, actual.Status)
		assert.Equal(t, "INVALID_BODY", actual.Errors[0].Flag)
		assert.NotEmpty(t, actual.Errors[0].Message)
		assert.Equal(t, "Validation error", actual.Message)
	})

	t.Run("it should return invalid body with invalid age format", func(t *testing.T) {
		defer db.DropAllTable()
		body := map[string]interface{}{
			"age": "36th",
		}
		w, _ := initTestUpdateUser("1", body)

		actual := response.Response{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, http.StatusBadRequest, actual.Status)
		assert.Equal(t, "INVALID_BODY", actual.Errors[0].Flag)
		assert.NotEmpty(t, actual.Errors[0].Message)
		assert.Equal(t, "Validation error", actual.Message)
	})

	t.Run("it should return invalid body with invalid param uri", func(t *testing.T) {
		defer db.DropAllTable()
		body := map[string]interface{}{
			"age": 10,
		}
		w, _ := initTestUpdateUser("x", body)

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
