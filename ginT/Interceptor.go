package ginT

import (
	"net/http"

	"github.com/ltto/T/ginT/ctx"
)

var (
	interceptors InterceptorList
)

type InterceptErr struct {
	Code  int
	Cause error
	isErr bool
}

func NewInterceptErr(code int, cause error) *InterceptErr {
	return &InterceptErr{Code: code, Cause: cause, isErr: cause != nil}
}

func NewInterceptOK() *InterceptErr {
	return &InterceptErr{Code: http.StatusOK, Cause: nil, isErr: false}
}

func (i InterceptErr) isError() bool {
	return i.isErr
}

func (i InterceptErr) Error() string {
	if i.Cause == nil {
		return "intercept Err Cause: nil"
	}
	return "intercept Err Cause: " + i.Cause.Error()
}

//-----------------------------------

type Interceptor func(*ctx.Context) *InterceptErr

type InterceptorList []Interceptor

func (i *InterceptorList) append(list ...Interceptor) {
	*i = append(*i, list...)
}

func AddInterceptor(list ...Interceptor) {
	interceptors.append(list...)
}

func DoInterceptorList(c *ctx.Context) *InterceptErr {
	if interceptors == nil {
		return NewInterceptOK()
	}
	var err *InterceptErr
	for _, v := range interceptors {
		if err = v(c); err.isErr {
			return err
		}
	}
	return err
}
