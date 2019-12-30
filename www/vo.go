package ct

import "github.com/go-sql-driver/mysql"

type Result struct {
	ErrNo  int         `json:"code"`
	ErrMsg string      `json:"errMsg"`
	Data   interface{} `json:"data"`
}

func Success(data interface{}) *Result {
	return &Result{0, "", data}
}

func Fail(errMsg string) *Result {
	return &Result{-1, errMsg, nil}
}
func Err(err error) *Result {
	return &Result{-1, err.Error(), nil}
}

func JudgeView(data interface{}, err error) *Result {
	if err != nil {
		return Fail(err.Error())
	} else {
		return Success(data)
	}
}

func errHandler(err error) string {
	if err != nil {
		if e, ok := err.(*mysql.MySQLError); ok {
			return e.Error()
		}
	}
	return ""
}
