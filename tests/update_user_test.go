package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
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

type updateUserSuccess struct {
	utils.Response
	Result repository.GetUser `json:"result"`
}

func initTestUpdateUser(id string, body map[string]interface{}) (*httptest.ResponseRecorder, *gin.Engine) {
	r := gin.Default()
	app := new(app.Application)
	app.CreateTest(r)

	w := httptest.NewRecorder()
	b, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPatch, "/user/"+id, strings.NewReader(string(b)))
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

func TestUpdateUserSuccess(t *testing.T) {
	defer db.DropAllTable()
	body := map[string]interface{}{
		"age": 35,
	}
	w, _ := initTestUpdateUser("1", body)

	actual := updateUserSuccess{}
	if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Success update user", actual.Message)
	assert.Equal(t, http.StatusOK, actual.Status)
	assert.NotEmpty(t, actual.Result)
}

func TestUpdateUserNotExist(t *testing.T) {
	defer db.DropAllTable()
	body := map[string]interface{}{
		"age": 35,
	}
	w, _ := initTestUpdateUser("2", body)

	actual := exception.Exception{}
	if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, http.StatusBadRequest, actual.Status)
	assert.Equal(t, "BAD_REQUEST", actual.Flag)
	assert.NotEmpty(t, actual.Errors)
	assert.Equal(t, "USER_NOT_FOUND", actual.Errors.Flag)
	assert.Equal(t, "User with this id not exist", actual.Errors.Message)
}

func TestUpdateUserInvalidBodyEmail(t *testing.T) {
	defer db.DropAllTable()
	body := map[string]interface{}{
		"email": "bukanatemaildotcom",
	}
	w, _ := initTestUpdateUser("1", body)

	actual := exception.Exception{}
	if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "BAD_REQUEST", actual.Flag)
	assert.NotEmpty(t, actual.Errors)
	assert.Equal(t, "INVALID_BODY", actual.Errors.Flag)
	assert.NotEmpty(t, actual.Errors.Message)
}

func TestUpdateUserInvalidBodyAge(t *testing.T) {
	defer db.DropAllTable()
	body := map[string]interface{}{
		"age": "36th",
	}
	w, _ := initTestUpdateUser("1", body)

	actual := exception.Exception{}
	if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "BAD_REQUEST", actual.Flag)
	assert.NotEmpty(t, actual.Errors)
	assert.Equal(t, "INVALID_BODY", actual.Errors.Flag)
	assert.NotEmpty(t, actual.Errors.Message)
}

func TestUpdateUserInvalidBodyParam(t *testing.T) {
	defer db.DropAllTable()
	body := map[string]interface{}{
		"age": 10,
	}
	w, _ := initTestUpdateUser("x", body)

	actual := exception.Exception{}
	if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "BAD_REQUEST", actual.Flag)
	assert.NotEmpty(t, actual.Errors)
	assert.Equal(t, "INVALID_BODY", actual.Errors.Flag)
	assert.NotEmpty(t, actual.Errors.Message)
}
