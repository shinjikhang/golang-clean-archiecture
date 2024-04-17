package common

import "net/http"

type successResponse struct {
	Status     int         `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination,omitempty"`
	Filter     interface{} `json:"filter,omitempty"`
}

func NewSuccessResponse(data interface{}, pagination interface{}, filter interface{}, message string) *successResponse {
	return &successResponse{
		Status:     http.StatusOK,
		Message:    message,
		Data:       data,
		Pagination: pagination,
		Filter:     filter,
	}
}

func SimpleSuccessResponse(data interface{}) *successResponse {
	return &successResponse{
		Data:       data,
		Pagination: nil,
		Filter:     nil,
		Status:     200,
		Message:    "success",
	}
}
