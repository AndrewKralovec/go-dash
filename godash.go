package godash

import (
	"reflect"
)

type ConditionalCallback func(v interface{}) bool

func Each(data interface{}, cb func(k interface{}, v interface{})) {
	v := reflect.ValueOf(data)

	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			val := v.Index(i).Interface()
			cb(i, val)
		}
	case reflect.Struct, reflect.Map:
		for i := 0; i < v.NumField(); i++ {
			val := v.Field(i).Interface()
			name := v.Type().Field(i).Name
			cb(name, val)
		}
	}
}

func Filter(data interface{}, cb ConditionalCallback) []interface{} {
	var array []interface{}
	v := reflect.ValueOf(data)
	for i := 0; i < v.Len(); i++ {
		val := v.Index(i).Interface()
		if cb(val) {
			array = append(array, val)
		}
	}

	return array
}

func Find(data interface{}, cb ConditionalCallback) interface{} {
	v := reflect.ValueOf(data)
	for i := 0; i < v.Len(); i++ {
		val := v.Index(i).Interface()
		if cb(val) {
			return val
		}
	}
	return nil
}

func First(data interface{}) interface{} {
	v := reflect.ValueOf(data)
	if v.Len() < 1 {
		return nil
	}

	val := v.Index(0).Interface()
	return val
}

func Includes(data interface{}, target interface{}) bool {
	v := reflect.ValueOf(data)
	for i := 0; i < v.Len(); i++ {
		val := v.Index(i).Interface()
		if val == target {
			return true
		}
	}
	return false
}

func Last(data interface{}) interface{} {
	v := reflect.ValueOf(data)
	val := v.Index(v.Len() - 1).Interface()
	return val
}
