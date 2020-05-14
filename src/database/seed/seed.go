package seed

import "github.com/jinzhu/gorm"

// Seed is seeding database configuration
type Seed struct {
	Name string
	Run  func(*gorm.DB) error
}

// All is running all seeder
func All() []Seed {
	return []Seed{
		{
			Name: "CreateUser",
			Run: func(db *gorm.DB) error {
				var err error
				for i := 0; i < 10; i++ {
					err = CreateUser(db)
					if err != nil {
						break
					}
				}
				return err
			},
		},
	}
}
