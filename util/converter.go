package util

import "reflect"

// ToSliceInterface converts a slice of type T to slice of interface.
// It expects a slice otherwise it'll return nil
func ToSliceInterface(s interface{}) []interface{} {
	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Slice {
		return nil
	}
	var slice []interface{}
	v := reflect.ValueOf(s)
	for i := 0; i < v.Len(); i++ {
		slice = append(slice, v.Index(i).Interface())
	}
	return slice
}
