package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/saefullohmaslul/golang-example/src/app"
	db "github.com/saefullohmaslul/golang-example/src/database"
	"github.com/saefullohmaslul/golang-example/src/database/entity"
	"github.com/saefullohmaslul/golang-example/src/repository"
	"github.com/saefullohmaslul/golang-example/src/utils"
	"github.com/stretchr/testify/assert"
)

type getUser struct {
	utils.Response
	Result repository.GetUser `json:"result"`
}

type getUserEmpty struct {
	utils.Response
	Errors getUserErrorMessage `json:"errors"`
}

type getUserError struct {
	Status int                 `json:"status"`
	Flag   string              `json:"flag"`
	Errors getUserErrorMessage `json:"errors"`
}

type getUserErrorMessage struct {
	Message string `json:"message"`
	Flag    string `json:"flag"`
}

func initTestGetUser(id string) (*httptest.ResponseRecorder, *gin.Engine) {
	r := gin.Default()
	app := new(app.Application)
	app.CreateTest(r)

	userRepository := repository.UserRepository{}
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

func TestGetUserSuccess(t *testing.T) {
	defer db.DropAllTable()
	w, _ := initTestGetUser("1")
	actual := getUser{}
	if err := json.Unmarshal([]byte(w.Body.String()), &actual); err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Success get user", actual.Message)
	assert.Equal(t, http.StatusOK, actual.Status)
	assert.NotEmpty(t, actual.Result)
}

func TestGetUserNotFound(t *testing.T) {
	defer db.DropAllTable()
	w, _ := initTestGetUser("2")
	actual := getUserEmpty{}
	if err := json.Unmarshal([]byte(w.Body.String()), &actual); err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "User not found", actual.Message)
	assert.Equal(t, http.StatusOK, actual.Status)
	assert.Equal(t, "User with this ID not enough", actual.Errors.Message)
	assert.Equal(t, "USER_NOT_FOUND", actual.Errors.Flag)
}

func TestGetUserErrorParamID(t *testing.T) {
	defer db.DropAllTable()
	w, _ := initTestGetUser("c")
	actual := getUserError{}
	if err := json.Unmarshal([]byte(w.Body.String()), &actual); err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Param must be of type integer, required", actual.Errors.Message)
	assert.Equal(t, http.StatusBadRequest, actual.Status)
	assert.Equal(t, "INVALID_BODY", actual.Errors.Flag)
	assert.Equal(t, "BAD_REQUEST", actual.Flag)
}
