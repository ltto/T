package param

type TT string

//integer
//number
//boolean
//string
//array
//object

type FMT string

//binary
//byte
//date
//date-time
//double
//email
//float
//int32
//int64
//uuid

type Param struct {
	//base
	Type   string `json:"type,omitempty"`
	Format string `json:"format,omitempty"`
	Schema Schema `json:"schema,omitempty"`
	//array
	Items *Param `json:"items,omitempty"`
	Ref   string `json:"$ref,omitempty"`
	//Obj
	//
	Default interface{} `json:"default,omitempty"`
}

func NewBaseParam() *Param {
	return &Param{}
}
func NewObjParam() *Param {
	return &Param{}
}
func NewArrParam(elem *Param) *Param {
	return &Param{Type: "array", Items: elem,}
}

//schema
type Schema struct {
	Ref string `json:"$ref"`
}
