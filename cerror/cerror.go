//author: wongoo
//date: 20190826

package cerror

import "net/http"

const (
	CodeOK         = 0
	CodeUnknownErr = 10
	CodeAuthErr    = 20
	CodeRequestErr = 100
)

var (
	ErrBadRequest   = NewStatusCodeError(http.StatusBadRequest, CodeRequestErr, "forbidden")
	ErrNotFound     = NewStatusCodeError(http.StatusNotFound, CodeRequestErr+1, "not found")
	ErrArgRequired  = NewStatusCodeError(http.StatusBadRequest, CodeRequestErr+2, "arg required")
	ErrValueInvalid = NewStatusCodeError(http.StatusBadRequest, CodeRequestErr+3, "value invalid")
	ErrUnauthorized = NewStatusCodeError(http.StatusUnauthorized, CodeAuthErr, "unauthorized")
	ErrForbidden    = NewStatusCodeError(http.StatusForbidden, CodeAuthErr+1, "forbidden")
)

type Coder interface {
	Code() int
}

type StatusState interface {
	Status() int
}

type CodeError interface {
	error
	Coder
}
type StatusCodeError interface {
	CodeError
	StatusState
}

func NewCodeError(code int, err string) CodeError {
	return &codeError{c: code, m: err}
}

type codeError struct {
	c int
	m string
}

func (e *codeError) Code() int {
	return e.c
}
func (e *codeError) Error() string {
	return e.m
}

func NewStatusCodeError(status, code int, err string) CodeError {
	return &statusCodeError{s: status, c: code, m: err}
}

type statusCodeError struct {
	s int
	c int
	m string
}

func (e *statusCodeError) Code() int {
	return e.c
}

func (e *statusCodeError) Status() int {
	return e.s
}

func (e *statusCodeError) Error() string {
	return e.m
}
