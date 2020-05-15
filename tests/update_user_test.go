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
	"github.com/saefullohmaslul/golang-example/src/database/entity"
	"github.com/saefullohmaslul/golang-example/src/middlewares/exception"
	"github.com/saefullohmaslul/golang-example/src/repositories"
	"github.com/saefullohmaslul/golang-example/src/utils"
	"github.com/saefullohmaslul/golang-example/src/utils/flag"
	"github.com/stretchr/testify/assert"
)

type updateUserSuccess struct {
	utils.Response
	Result repositories.GetUser `json:"result"`
}

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

		actual := updateUserSuccess{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, flag.UpdateUserSuccess.Message, actual.Message)
		assert.Equal(t, http.StatusOK, actual.Status)
		assert.NotEmpty(t, actual.Result)
	})

	t.Run("it should return user not exist", func(t *testing.T) {
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
		assert.Equal(t, flag.UpdateUserNotExist.Flag, actual.Flag)
		assert.NotEmpty(t, actual.Errors)
		assert.Equal(t, flag.UpdateUserNotExist.Error.Flag, actual.Errors.Flag)
		assert.Equal(t, flag.UpdateUserNotExist.Error.Message, actual.Errors.Message)
	})

	t.Run("it should return invalid body with invalid email format", func(t *testing.T) {
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
		assert.Equal(t, flag.UpdateUserInvalidBody.Flag, actual.Flag)
		assert.NotEmpty(t, actual.Errors)
		assert.Equal(t, flag.UpdateUserInvalidBody.Error.Flag, actual.Errors.Flag)
		assert.NotEmpty(t, actual.Errors.Message)
	})

	t.Run("it should return invalid body with invalid age format", func(t *testing.T) {
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
		assert.Equal(t, flag.UpdateUserInvalidBody.Flag, actual.Flag)
		assert.NotEmpty(t, actual.Errors)
		assert.Equal(t, flag.UpdateUserInvalidBody.Error.Flag, actual.Errors.Flag)
		assert.NotEmpty(t, actual.Errors.Message)
	})

	t.Run("it should return invalid body with invalid param uri", func(t *testing.T) {
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
		assert.Equal(t, flag.UpdateUserInvlidParamURI.Flag, actual.Flag)
		assert.NotEmpty(t, actual.Errors)
		assert.Equal(t, flag.UpdateUserInvlidParamURI.Error.Flag, actual.Errors.Flag)
		assert.Equal(t, flag.UpdateUserInvlidParamURI.Error.Message, actual.Errors.Message)
	})
}
