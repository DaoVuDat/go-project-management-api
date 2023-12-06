package domain

import (
	"errors"
)

var (
	ErrUsernameExists       = errors.New("username already exists")
	ErrBadRequest           = errors.New("bad request")
	ErrInvalidUserAccountId = errors.New("invalid user account id")
)

type ErrResponse struct {
	HttpStatusCode int    `json:"code"`
	Message        string `json:"message"`
	Detail         string `json:"detail"`
}

func ErrInvalidRequest(err error) *ErrResponse {
	return &ErrResponse{
		HttpStatusCode: 400,
		Message:        "Invalid request.",
		Detail:         err.Error(),
	}
}

func ErrRender(err error) *ErrResponse {
	return &ErrResponse{
		HttpStatusCode: 422,
		Message:        "Error rendering response.",
		Detail:         err.Error(),
	}
}

func ErrInternal(err error) *ErrResponse {
	return &ErrResponse{
		HttpStatusCode: 500,
		Message:        "Internal server error.",
		Detail:         err.Error(),
	}
}

func ErrResourceConflict(err error) *ErrResponse {
	return &ErrResponse{
		HttpStatusCode: 409,
		Message:        "Internal server error.",
		Detail:         err.Error(),
	}
}

var ErrNotFound = &ErrResponse{HttpStatusCode: 404, Message: "Resource not found."}
