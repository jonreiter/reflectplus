package reflectplus

import (
	"log"
	"reflect"
)

// LookLikeAliases returns true when these two objects look like aliases
// this isn't something we can determine for sure at runtime
func LookLikeAliases(a, b interface{}) bool {
	aValue := reflect.ValueOf(a).Elem()
	bValue := reflect.ValueOf(b).Elem()
	bValueType := bValue.Type()
	aValueType := aValue.Type()
	nFields := bValueType.NumField()
	if nFields != aValueType.NumField() {
		log.Panic("aliasCopy nFields mismatch")
	}
	for i := 0; i < nFields; i++ {
		if bValueType.Field(i).Name != aValueType.Field(i).Name {
			log.Panic("aliasCopy field name mismatch")
		}
		if bValueType.Field(i).Type != aValueType.Field(i).Type {
			log.Panic("aliasCopy field type mismatch")
		}
	}
	return true
}

// IsPointer returns true if the passed in object is a pointer
func IsPointer(a interface{}) bool {
	aType := reflect.TypeOf(a)
	aKind := aType.Kind()
	if aKind != reflect.Ptr {
		return false
	}
	return true
}

// AliasCopy copies everything from from into to
// it requires that to and from LookLikeAliases()
// and that both are pointers
func AliasCopy(to, from interface{}) {
	if !IsPointer(to) {
		panic("trying to aliasCopy to isn't a pointer")
	}
	if !IsPointer(from) {
		panic("trying to aliasCopy from isn't a pointer")
	}
	if !LookLikeAliases(to, from) {
		panic("trying to aliasCopy objects that do not look like aliases")
	}
	fromValue := reflect.ValueOf(from).Elem()
	toValue := reflect.ValueOf(to).Elem()
	toValueType := toValue.Type()
	nFields := toValueType.NumField()
	for i := 0; i < nFields; i++ {
		toValueField := toValue.Field(i)
		theValue := fromValue.FieldByName(toValueType.Field(i).Name)
		if toValueField.CanSet() {
			toValueField.Set(theValue)
		}
	}
}

// eof
