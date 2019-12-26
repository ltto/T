package vo

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Page *Page       `json:"page"`
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

type Page struct {
	Page  int `query:"page"`
	Size  int `query:"size"`
	Count int
}

func (v *Page) Limit() (int, int) {
	if v.Page <= 0 {
		v.Page = 1
	}
	if v.Size <= 0 {
		v.Size = 8
	}
	return (v.Page - 1) * v.Size, v.Size
}
