package types

import "github.com/saefullohmaslul/Golang-Example/utils"

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
