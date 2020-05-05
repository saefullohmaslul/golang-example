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

func initTestGetUsers() (*httptest.ResponseRecorder, *gin.Engine) {
	r := gin.Default()
	app := new(app.Application)
	app.CreateTest(r)

	userRepository := repository.UserRepository{}
	userRepository.CreateUser(entity.User{
		ID:       1,
		Address:  "Jakarta",
		Age:      20,
		Email:    "email1@email.com",
		Name:     "Saefulloh Maslul",
		Password: "123456",
	})
	userRepository.CreateUser(entity.User{
		ID:       2,
		Address:  "Jakarta",
		Age:      20,
		Email:    "email2@email.com",
		Name:     "Saefulloh Maslul",
		Password: "123456",
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)
	return w, r
}

type getUsers struct {
	utils.Response
	Result []repository.GetUser `json:"result"`
}

func TestGetUsersSuccess(t *testing.T) {
	defer db.DropAllTable()
	w, _ := initTestGetUsers()
	actual := getUsers{}
	if err := json.Unmarshal([]byte(w.Body.String()), &actual); err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Success get all users", actual.Message)
	assert.Equal(t, http.StatusOK, actual.Status)
	assert.NotEmpty(t, actual.Result)
}
