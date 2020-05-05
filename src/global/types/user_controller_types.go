package types

import (
	"github.com/saefullohmaslul/golang-example/src/repository"
	"github.com/saefullohmaslul/golang-example/src/utils"
)

// GetNameResponse is return to get name format
type GetNameResponse struct {
	utils.Response
	Result string `json:"result"`
}

// GetBiodataResult is result format to GetBiodataResponse
type GetBiodataResult struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

// GetBiodataResponse is return to get biodata format
type GetBiodataResponse struct {
	utils.Response
	Result GetBiodataResult `json:"result"`
}

// CreateUserResult format
type CreateUserResult struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Age     int64  `json:"age"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

// CreateUserResponse is return to create user into database
type CreateUserResponse struct {
	utils.Response
	Result repository.GetUser `json:"result"`
}
