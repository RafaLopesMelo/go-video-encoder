package httperror

import (
	"net/http"

	domainerrors "github.com/RafaLopesMelo/go-video-encoder/internal/domain/errors"
)

type HttpError struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Doc     string `json:"documentation"`
}

func new(status int, code string, message string, doc string) HttpError {
	return HttpError{
		Status:  status,
		Code:    code,
		Message: message,
		Doc:     doc,
	}
}

func NewFromDomain(err domainerrors.DomainError) HttpError {
	missingCode := "0000"
	missingDoc := ""

	switch err {
	case domainerrors.RequiredProperty:
		return new(http.StatusUnprocessableEntity, missingCode, err.Error(), missingDoc)
	case domainerrors.EntityNotFound:
		return new(http.StatusNotFound, missingCode, err.Error(), missingDoc)
	default:
		return new(http.StatusInternalServerError, missingCode, "Internal Server Error", missingDoc)
	}
}
