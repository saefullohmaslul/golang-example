package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/saefullohmaslul/golang-example/src/apps"
	db "github.com/saefullohmaslul/golang-example/src/database"
	"github.com/saefullohmaslul/golang-example/src/database/entity"
	"github.com/saefullohmaslul/golang-example/src/repositories"
	"github.com/saefullohmaslul/golang-example/src/utils/flag"
	"github.com/saefullohmaslul/golang-example/src/utils/response"
	"github.com/stretchr/testify/assert"
)

func initTestGetUser(id string) (*httptest.ResponseRecorder, *gin.Engine) {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	userRepository := repositories.UserRepository{Conn: db.GetDB().Table("users")}
	userRepository.CreateUser(entity.User{
		ID:       1,
		Address:  "Jakarta",
		Age:      20,
		Email:    "email@email.com",
		Name:     "Saefulloh Maslul",
		Password: "123456",
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user/"+id, nil)
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)
	return w, r
}

func TestGetUser(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer db.DropAllTable()
		w, _ := initTestGetUser("1")
		actual := response.Response{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, http.StatusOK, actual.Status)
		assert.Equal(t, flag.GetUserSuccess.Message, actual.Message)
		assert.NotEmpty(t, actual.Data)
		assert.Empty(t, actual.Errors)
	})

	t.Run("it should return user not found", func(t *testing.T) {
		defer db.DropAllTable()
		w, _ := initTestGetUser("2")
		actual := response.Response{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, http.StatusNotFound, actual.Status)
		assert.Equal(t, "User not found", actual.Message)
		assert.Equal(t, "User with this ID not found", actual.Errors[0].Message)
		assert.Equal(t, "USER_NOT_FOUND", actual.Errors[0].Flag)
	})

	t.Run("it should return invalid body with invalid param uri", func(t *testing.T) {
		defer db.DropAllTable()
		w, _ := initTestGetUser("c")
		actual := response.Response{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, http.StatusBadRequest, actual.Status)
		assert.Equal(t, "Validation error", actual.Message)
		assert.Equal(t, "Param must be of type integer, required", actual.Errors[0].Message)
		assert.Equal(t, "INVALID_BODY", actual.Errors[0].Flag)
	})
}
