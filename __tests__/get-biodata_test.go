package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/saefullohmaslul/Golang-Example/controllers"
	"github.com/stretchr/testify/assert"
)

func TestGetBiodata(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/user/biodata", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user/biodata")

	if assert.NoError(t, controllers.GetBiodata(c)) {
		expect := `{"status":200,"message":"Success to get biodata","result": {"name": "Saefulloh Maslul","address": "Tegal"}}`

		res := controllers.GetBiodataResponse{}
		mock := controllers.GetBiodataResponse{}

		if err := json.Unmarshal([]byte(rec.Body.String()), &res); err != nil {
			panic(err)
		}

		if err := json.Unmarshal([]byte(expect), &mock); err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, mock.Message, res.Message)
		assert.Equal(t, mock.Status, res.Status)
		assert.NotZero(t, res.Result)
	}
}
