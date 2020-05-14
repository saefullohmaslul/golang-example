package seed

import (
	"github.com/bxcodec/faker"
	"github.com/jinzhu/gorm"
	"github.com/saefullohmaslul/golang-example/src/database/entity"
	"github.com/sirupsen/logrus"
)

// CreateUser is seeder to create user
func CreateUser(db *gorm.DB) error {
	user := entity.User{}
	if err := faker.FakeData(&user); err != nil {
		logrus.Errorln("Error user seed", err)
	}

	return db.Create(&entity.User{
		Name:     user.Name,
		Address:  user.Address,
		Age:      user.Age,
		Email:    user.Email,
		Password: user.Password,
	}).Error
}
