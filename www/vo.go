package www

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
