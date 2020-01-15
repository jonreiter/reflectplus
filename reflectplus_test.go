package reflectplus_test

import (
	"reflect"
	"testing"

	"github.com/jonreiter/reflectplus"
)

type CopyStruct struct {
	A int
	B float64
}

func TestDuplicate(t *testing.T) {
	i := CopyStruct{A: 10, B: 1.234}
	var j CopyStruct

	reflectplus.AliasCopy(&j, &i)
	if i.A != j.A || i.B != j.B {
		t.Error("copy problem")
	}
	newObj := reflect.New(reflect.TypeOf(i)).Interface()
	newObjS := newObj.(*CopyStruct)
	reflectplus.AliasCopy(newObjS, &i)
	if i.A != newObjS.A || i.B != newObjS.B {
		t.Error("copy problem")
	}
}

// eof
