package entity

import "time"

// User -> user entity schema
type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Name      string     `json:"name" faker:"name"`
	Age       int64      `json:"age"`
	Email     string     `gorm:"type:varchar(100);unique_index" json:"email" faker:"email"`
	Address   string     `gorm:"index:addr" json:"address" faker:"word"`
	Password  string     `gorm:"type:varchar(255)" json:"password" faker:"password"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
