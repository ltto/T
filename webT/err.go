package webT

import (
	"fmt"
)

type HTTPError struct {
	Code int
	Msg  string
}

func (H HTTPError) Error() string {
	return fmt.Sprintf("code:%d msg:%s", H.Code, H.Msg)
}

func NewHTTPError(code int, msg string) *HTTPError {
	return &HTTPError{Code: code, Msg: msg}
}
