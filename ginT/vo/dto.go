package vo

type Page struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Count int `json:"count"`
}

func NewPage(page int, size int) *Page {
	v := Page{Page: page, Size: size}
	if v.Page <= 0 {
		v.Page = 1
	}
	if v.Size <= 0 {
		v.Size = 8
	}
	return &v
}

func (v Page) Limit() (int, int) {
	return (v.Page - 1) * v.Size, v.Size
}
