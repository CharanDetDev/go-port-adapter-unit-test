package model

type (
	Response struct {
		Code       string      `json:"code"`
		Data       interface{} `json:"data,omitempty"`
		HTTPStatus int         `json:"httpStatus,omitempty"`
		Message    string      `json:"message"`
		Errors     interface{} `json:"errors,omitempty"`
	}

	ErrorResponseModel struct {
		Code       string      `json:"code"`
		HTTPStatus int         `json:"httpStatus,omitempty"`
		Message    string      `json:"message"`
		Errors     interface{} `json:"errors,omitempty"`
	}
)
