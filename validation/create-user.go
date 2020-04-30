package validation

// CreateUserSchema -> create user schema validation
type CreateUserSchema struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Age      int64
	Address  string
}
