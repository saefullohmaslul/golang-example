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

func TestGetBiodata(t *testing.T) {
	r := gin.Default()
	app := new(app.Application)
	app.CreateApp(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user/biodata", nil)
	r.ServeHTTP(w, req)

	e := `{"status":200,"message":"Success get biodata","result": {"name": "Saefulloh Maslul","address": "Tegal"}}`

	actual := types.GetBiodataResponse{}
	expect := types.GetBiodataResponse{}

	if err := json.Unmarshal([]byte(w.Body.String()), &actual); err != nil {
		panic(err)
	}
	if err := json.Unmarshal([]byte(e), &expect); err != nil {
		panic(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expect.Message, actual.Message)
	assert.Equal(t, expect.Status, actual.Status)
	assert.Equal(t, expect.Result.Name, actual.Result.Name)
	assert.Equal(t, expect.Result.Address, actual.Result.Address)
}
