package handlerutil

import (
	"errors"
	"net/http"
)

type HttpError interface {
	Error() string
	StatusCode() int
}

type badRequest struct {
	error
}

func BadRequest(err string) *badRequest {
	return &badRequest{errors.New(err)}
}

func (e *badRequest) StatusCode() int {
	return http.StatusBadRequest
}

type notFound struct {
	error
}

func NotFound(err string) *notFound {
	return &notFound{errors.New(err)}
}

func (e *notFound) StatusCode() int {
	return http.StatusNotFound
}

type badGateway struct {
	error
}

func BadGateway(err string) *badGateway {
	return &badGateway{errors.New(err)}
}

func (e *badGateway) StatusCode() int {
	return http.StatusBadGateway
}

type unauthorized struct {
	error
}

func Unauthorized(err string) *unauthorized {
	return &unauthorized{errors.New(err)}
}

func (e *unauthorized) StatusCode() int {
	return http.StatusUnauthorized
}

type forbidden struct {
	error
}

func Forbidden(err string) *forbidden {
	return &forbidden{errors.New(err)}
}

func (e *forbidden) StatusCode() int {
	return http.StatusForbidden
}

type tooManyRequests struct {
	error
}

func TooManyRequests(err string) *tooManyRequests {
	return &tooManyRequests{errors.New(err)}
}

func (e *tooManyRequests) StatusCode() int {
	return http.StatusTooManyRequests
}

type internalServerError struct {
	error
}

func InternalServerError(err string) *internalServerError {
	return &internalServerError{errors.New(err)}
}

func (e *internalServerError) StatusCode() int {
	return http.StatusInternalServerError
}

func GenericServerError() *internalServerError {
	return InternalServerError("something went wrong")
}
