package services

import (
	"restapi/src/repositories"
	"testing"

	"github.com/sarulabs/di"
	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	assert := assert.New(t)

	t.Run("It should passing ios into service", func(t *testing.T) {
		builder, _ := di.NewBuilder()
		builder.Add(di.Def{
			Name: "repository",
			Build: func(ctn di.Container) (interface{}, error) {
				return repositories.NewRepositoryMock(builder.Build()), nil
			},
		})

		service := NewService(builder.Build())
		assert.NotNil(t, service)
	})
}
