package utils

import (
	"reflect"
	"strconv"
)

func GetParams(object any, formatType string) map[string]any {
	var configMap = make(map[string]any)
	objType := reflect.TypeOf(object)
	objValue := reflect.ValueOf(object)

	for i := 0; i < objValue.NumField(); i++ {
		switch objValue.Field(i).Kind() {
		case reflect.Ptr:
			objType = objType.Elem()
			objValue = objValue.Elem()
			fallthrough
		case reflect.Struct:
			configMap[objType.Field(i).Tag.Get(formatType)] = GetParams(objValue.Field(i).Interface(), formatType)
		default:
			configMap[objType.Field(i).Tag.Get(formatType)] = objValue.Field(i)
		}
	}

	return configMap
}

func FieldType(object any, name string) reflect.Kind {
	params := reflect.ValueOf(object).Elem()
	fields := params.Type()
	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		if field.Name == name {
			return field.Type.Kind()
		}
	}
	return reflect.Invalid
}

func SetTagData(obj any, name string) {
	params := reflect.ValueOf(obj).Elem()
	fields := params.Type()
	for i := 0; i < fields.NumField(); i++ {
		param := params.Field(i)
		field := fields.Field(i)
		if value, ok := field.Tag.Lookup(name); ok {
			switch field.Type.Kind() {
			case reflect.String:
				if param.String() == "" {
					param.SetString(value)
				}
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if param.Int() == 0 {
					i, _ := strconv.Atoi(value)
					param.SetInt(int64(i))
				}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				if param.Uint() == 0 {
					i, _ := strconv.Atoi(value)
					param.SetUint(uint64(i))
				}
			case reflect.Float32, reflect.Float64:
				if param.Float() == 0 {
					i, _ := strconv.Atoi(value)
					param.SetFloat(float64(i))
				}
			case reflect.Bool:
				if param.Bool() == false {
					b, _ := strconv.ParseBool(value)
					param.SetBool(b)
				}
			default:
				v := reflect.ValueOf(value)
				param.Set(v)
			}
		}
	}
}

func TagValue(obj any, tag, name string) string {
	var tagValue string
	params := reflect.ValueOf(obj).Elem()
	fields := params.Type()
	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		if field.Name == name {
			if value, ok := field.Tag.Lookup(tag); ok {
				tagValue = value
				break
			}
		}
	}
	return tagValue
}

func FieldByTagValue(obj any, tag, value string) string {
	var fieldName string
	params := reflect.ValueOf(obj).Elem()
	fields := params.Type()
	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		if v, ok := field.Tag.Lookup(tag); ok {
			if v == value {
				fieldName = field.Name
				break
			}
		}
	}
	return fieldName
}

func TypeValue(obj any, tag, value string) string {
	var tagValue string
	params := reflect.ValueOf(obj).Elem()
	fields := params.Type()
	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		if field.Type.Name() == value {
			if v, ok := field.Tag.Lookup(tag); ok {
				tagValue = v
				break
			}
		}
	}
	return tagValue
}

func SetDefaults(obj any) {
	SetTagData(obj, "default")
}

func DefaultValue(obj interface{}, name string) string {
	return TagValue(obj, "default", name)
}

func TypeDefaultValue(obj interface{}, name string) string {
	return TypeValue(obj, "default", name)
}
