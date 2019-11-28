package ref

import (
	"testing"
)

/**func TestClonePtr(t *testing.T) {
	var s int
	//value := NewSetV(reflect.TypeOf(s), reflect.ValueOf(50))
	value := NewSetV(reflect.TypeOf(s), 1000)
	if value.Interface() != 1000 {
		t.Error("set err")
	}
}**/

func TestSrcTypeVal(t *testing.T) {
	var i ***int
	value, i2 := SrcTypeVal(i)
	t.Log(value, i2)
}
