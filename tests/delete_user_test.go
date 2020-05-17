package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
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

func initTestDeleteUser(id string) (*httptest.ResponseRecorder, *gin.Engine) {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/user/"+id, nil)
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

func TestDeleteUser(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer db.DropAllTable()
		w, _ := initTestDeleteUser("1")

		actual := response.Response{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, http.StatusOK, actual.Status)
		assert.Equal(t, flag.DeleteUserSuccess.Message, actual.Message)
		assert.NotEmpty(t, actual.Data)
	})

	t.Run("it should return user not found", func(t *testing.T) {
		defer db.DropAllTable()
		w, _ := initTestDeleteUser("2")

		actual := response.Response{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, http.StatusNotFound, actual.Status)
		assert.Equal(t, "User not exist", actual.Message)
		assert.Equal(t, "User with this ID not found", actual.Errors[0].Message)
		assert.Equal(t, "USER_NOT_FOUND", actual.Errors[0].Flag)
	})

	t.Run("it should return invalid param uri with invalid id format", func(t *testing.T) {
		defer db.DropAllTable()
		w, _ := initTestDeleteUser("x")

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
