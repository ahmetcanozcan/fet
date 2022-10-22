// Package fet implements StructUpdater that allows to building filters using a struct directly
package fet

import (
	"errors"
	"reflect"
)

var (
	ErrNilStruct    = errors.New("fet: nil pointer passed to Struct")
	ErrNotStruct    = errors.New("fet: non-struct passed to Struct")
	ErrInvalidField = errors.New("fet: invalid field")
)

// StructUpdater is and updater for struct
type StructUpdater struct {
	v reflect.Value
	t reflect.Type
	m M
}

// Pick picks the given field from struct with the given checkers.
// if the struct field has `bson` tag, otherwise use fieldname
// Example:
// 	fet.Struct(user).Pick("firstname").Pick("lastname")
func (s *StructUpdater) Pick(name string, checkers ...Checker) *StructUpdater {
	vf := s.v.FieldByName(name)

	f, ok := s.t.FieldByName(name)

	if !ok {
		return s
	}

	val := vf.Interface()

	key := getKeyFromField(f, name)

	if !checkFilter(key, val, checkers...) {
		return s
	}

	s.m[key] = val

	return s
}

// PickAll picks all the fields from the given struct.
// use given checkers to all fields.
func (s *StructUpdater) PickAll(checkers ...Checker) *StructUpdater {
	for i := 0; i < s.t.NumField(); i++ {
		s.Pick(s.t.Field(i).Name, checkers...)
	}

	return s
}

// Update applies all picked struct fields to given map
func (s *StructUpdater) Update(m M) {
	for k, v := range s.m {
		m[k] = v
	}
}

// Build builds a fet.M from picked struct fields, equivalent of:
// 		Build(structUpdater)
func (s *StructUpdater) Build() M {
	return Build(s)
}

// Struct constructs a StructUpdater for the given struct.
// if given value is nil or not a struct, panic.
func Struct(s interface{}) *StructUpdater {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	if s == nil {
		panic(ErrNilStruct)
	}

	kind := v.Kind()

	if kind == reflect.Ptr {
		return Struct(v.Elem().Interface())
	}

	if kind != reflect.Ptr && kind != reflect.Struct {
		panic(ErrNotStruct)
	}

	return &StructUpdater{
		v: v,
		t: t,
		m: M{},
	}
}
