package godash

import "reflect"

func First(data interface{}) interface{} {
	v := reflect.ValueOf(data)
	if v.Len() < 1 {
		return nil
	}

	val := v.Index(0).Interface()
	return val
}