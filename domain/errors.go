package domain

import (
	"errors"
	"net/http"
)

var (
	ErrUsernameExists       = errors.New("username already exists")
	ErrBadRequest           = errors.New("bad request")
	ErrInvalidUserAccountId = errors.New("invalid user account id")
	ErrInvalidLogin         = errors.New("username does not exists or password incorrect")
	ErrInvalidToken         = errors.New("invalid token")
	ErrExpiredToken         = errors.New("expired token")
	ErrInvalidAuthorization = errors.New("invalid authorization")
)

func ErrCustom(errorStr string) error {
	return errors.New(errorStr)
}

type ErrResponse struct {
	HttpStatusCode int    `json:"code"`
	Message        string `json:"message"`
	Detail         string `json:"detail"`
}

func ErrInvalidRequestResponse(err error) *ErrResponse {
	return &ErrResponse{
		HttpStatusCode: http.StatusBadRequest,
		Message:        "Invalid request.",
		Detail:         err.Error(),
	}
}

func ErrRenderResponse(err error) *ErrResponse {
	return &ErrResponse{
		HttpStatusCode: http.StatusUnprocessableEntity,
		Message:        "Error rendering response.",
		Detail:         err.Error(),
	}
}

func ErrInternalResponse(err error) *ErrResponse {
	return &ErrResponse{
		HttpStatusCode: http.StatusInternalServerError,
		Message:        "Internal server error.",
		Detail:         err.Error(),
	}
}

func ErrResourceConflictResponse(err error) *ErrResponse {
	return &ErrResponse{
		HttpStatusCode: http.StatusConflict,
		Message:        "Internal server error.",
		Detail:         err.Error(),
	}
}

func ErrInvalidLoginResponse(err error) *ErrResponse {
	return &ErrResponse{
		HttpStatusCode: http.StatusUnauthorized,
		Message:        "Unauthorized",
		Detail:         err.Error(),
	}
}

func ErrUnauthorizedResponse(err error) *ErrResponse {
	return &ErrResponse{
		HttpStatusCode: http.StatusUnauthorized,
		Message:        "Unauthorized",
		Detail:         err.Error(),
	}
}

var ErrNotFound = &ErrResponse{HttpStatusCode: 404, Message: "Resource not found."}
