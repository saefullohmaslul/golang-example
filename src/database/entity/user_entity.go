package entity

// User -> user entity schema
type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Name     string `json:"name"`
	Age      int64  `json:"age"`
	Email    string `gorm:"type:varchar(100);unique_index" json:"email"`
	Address  string `gorm:"index:addr" json:"address"`
	Password string `gorm:"type:varchar(255)" json:"password"`
}
