package validations_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/saefullohmaslul/golang-example/src/validations"
)

func TestValidate(t *testing.T) {
	var errors []map[string]interface{}
	schemas := "Hello World"
	defer func() {
		if err := recover(); err != nil {
			assert.NotEmpty(t, err)
		}
	}()
	validations.Validate(schemas, errors)
}
