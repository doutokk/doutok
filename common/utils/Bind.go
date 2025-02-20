package utils

import (
	"github.com/cloudwego/hertz/pkg/app"
	"reflect"
	"strconv"
	"strings"
)

// todo:性能？还是改rest
func BindParamsToStruct(c *app.RequestContext, req interface{}) error {
	val := reflect.ValueOf(req).Elem()
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		paramName := strings.ToLower(field.Name[:1]) + field.Name[1:]
		paramValue, exists := c.Params.Get(paramName)
		if !exists {
			continue
		}

		fieldValue := val.Field(i)
		if !fieldValue.CanSet() {
			continue
		}

		switch fieldValue.Kind() {
		case reflect.String:
			fieldValue.SetString(paramValue)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intValue, err := strconv.ParseInt(paramValue, 10, 64)
			if err != nil {
				return err
			}
			fieldValue.SetInt(intValue)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			uintValue, err := strconv.ParseUint(paramValue, 10, 64)
			if err != nil {
				return err
			}
			fieldValue.SetUint(uintValue)
		case reflect.Float32, reflect.Float64:
			floatValue, err := strconv.ParseFloat(paramValue, 64)
			if err != nil {
				return err
			}
			fieldValue.SetFloat(floatValue)
		}
	}
	return nil
}
