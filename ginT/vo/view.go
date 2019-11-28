package vo

type Result struct {
	ErrNo  int         `json:"code"`
	ErrMsg string      `json:"errMsg"`
	Data   interface{} `json:"data"`
	Page   *Page       `json:"page"`
}

func List(data interface{}, page *Page) *Result {
	return &Result{0, "", data, page}
}

func Success(data interface{}) *Result {
	return &Result{0, "", data, nil}
}

func Fail(errMsg string) *Result {
	return &Result{-1, errMsg, nil, nil}
}
func Err(err error) *Result {
	return &Result{-1, err.Error(), nil, nil}
}

func JudgeView(data interface{}, err error) *Result {
	if err != nil {
		return Fail(err.Error())
	} else {
		return Success(data)
	}
}
