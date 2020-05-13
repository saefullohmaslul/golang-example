package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/saefullohmaslul/golang-example/src/app"
	db "github.com/saefullohmaslul/golang-example/src/database"
	"github.com/saefullohmaslul/golang-example/src/database/entity"
	"github.com/saefullohmaslul/golang-example/src/middlewares/exception"
	"github.com/saefullohmaslul/golang-example/src/repository"
	"github.com/saefullohmaslul/golang-example/src/utils"
	"github.com/stretchr/testify/assert"
)

type deleteUserSuccess struct {
	utils.Response
	Result repository.GetUser `json:"result"`
}

func initTestDeleteUser(id string) (*httptest.ResponseRecorder, *gin.Engine) {
	r := gin.Default()
	app := new(app.Application)
	app.CreateTest(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/user/"+id, nil)
	req.Header.Set("Content-Type", "application/json")

	userRepository := repository.UserRepository{Conn: db.GetDB().Table("users")}
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

func TestDeleteUserSuccess(t *testing.T) {
	defer db.DropAllTable()
	w, _ := initTestDeleteUser("1")

	fmt.Println(w.Body.String())

	actual := deleteUserSuccess{}
	if err := json.Unmarshal([]byte(w.Body.String()), &actual); err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Success delete user", actual.Message)
	assert.Equal(t, http.StatusOK, actual.Status)
	assert.NotEmpty(t, actual.Result)
}

func TestDeleteUserNotExist(t *testing.T) {
	defer db.DropAllTable()
	w, _ := initTestDeleteUser("2")

	actual := exception.Exception{}
	if err := json.Unmarshal([]byte(w.Body.String()), &actual); err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, http.StatusBadRequest, actual.Status)
	assert.Equal(t, "BAD_REQUEST", actual.Flag)
	assert.NotEmpty(t, actual.Errors)
	assert.Equal(t, "USER_NOT_FOUND", actual.Errors.Flag)
	assert.Equal(t, "User with this id not exist", actual.Errors.Message)
}

func TestDeleteUserInvalidBodyParam(t *testing.T) {
	defer db.DropAllTable()
	w, _ := initTestDeleteUser("x")

	actual := exception.Exception{}
	if err := json.Unmarshal([]byte(w.Body.String()), &actual); err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "BAD_REQUEST", actual.Flag)
	assert.NotEmpty(t, actual.Errors)
	assert.Equal(t, "INVALID_BODY", actual.Errors.Flag)
	assert.NotEmpty(t, actual.Errors.Message)
}
