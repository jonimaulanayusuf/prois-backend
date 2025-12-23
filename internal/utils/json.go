package utils

import (
	"reflect"
)

func GetJSONFieldName(obj interface{}, fieldName string) string {
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if f, ok := t.FieldByName(fieldName); ok {
		jsonTag := f.Tag.Get("json")
		if jsonTag != "" {
			return jsonTag
		}
	}
	return fieldName
}
