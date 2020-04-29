package validation

// CreateUserSchema -> create user schema validation
type CreateUserSchema struct {
	Name string `validate:"required"`
}
