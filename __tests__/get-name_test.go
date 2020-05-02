package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/saefullohmaslul/Golang-Example/app"
	"github.com/saefullohmaslul/Golang-Example/global/types"
	"github.com/stretchr/testify/assert"
)

func TestGetName(t *testing.T) {
	r := gin.Default()
	app := new(app.Application)
	app.CreateApp(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/name", nil)
	r.ServeHTTP(w, req)

	actual := types.GetNameResponse{}

	if err := json.Unmarshal([]byte(w.Body.String()), &actual); err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Success get name", actual.Message)
	assert.Equal(t, http.StatusOK, actual.Status)
	assert.NotEmpty(t, actual.Result)
}
