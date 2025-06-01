package model

type ResponseJSON struct {
	Message         string      `json:"message"`
	Data            interface{} `json:"data,omitempty"`
	Success         string      `json:"success"`
	Error           string      `json:"error,omitempty"`
	ValidationError interface{} `json:"validation_error,omitempty"`
}